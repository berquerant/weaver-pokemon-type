package main_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	floatEqualTorelance = 0.0001
	endpoint            = "http://localhost:21099"
	weaverConfig        = "../../weaver.toml"
	serverPattern       = "."
)

func floatEqual(a, b float32) bool {
	return math.Abs(float64(a-b)) < floatEqualTorelance
}

func TestAPI(t *testing.T) {
	os.Setenv("SERVICEWEAVER_CONFIG", weaverConfig)
	stop, err := runMain(serverPattern)
	if err != nil {
		t.Fatal(err)
	}
	defer stop()

	client := &client{
		client: &http.Client{
			Timeout: time.Second,
		},
	}
	client.init(t)

	t.Run("type", func(t *testing.T) {
		for _, tc := range []typeTestcase{
			{
				title:   "くさ",
				request: []byte(`{"name":"くさ"}`),
				wantID:  5,
			},
			{
				title:   "ゴースト",
				request: []byte(`{"name":"ゴースト"}`),
				wantID:  14,
			},
		} {
			tc := tc
			t.Run(tc.title, func(t *testing.T) {
				tc.test(t, client)
			})
		}
	})

	t.Run("attack", func(t *testing.T) {
		for _, tc := range []attackTestcase{
			{
				title:   "エスパー2倍",
				request: []byte(`{"id":11,"index":1}`),
				filter:  func(item item) bool { return floatEqual(item.Multiplier, 2) },
				want: items{
					newItem(2, "エスパー", "かくとう"),
					newItem(2, "エスパー", "どく"),
				},
			},
			{
				title:   "あく0.25倍複合",
				request: []byte(`{"id":16,"index":2}`),
				filter:  func(item item) bool { return floatEqual(item.Multiplier, 0.25) },
				want: items{
					newItem(0.25, "あく", "あく", "あく", "かくとう"),
					newItem(0.25, "あく", "かくとう", "あく", "フェアリー"),
					newItem(0.25, "あく", "あく", "あく", "フェアリー"),
				},
			},
		} {
			tc := tc
			t.Run(tc.title, func(t *testing.T) {
				tc.test(t, client)
			})
		}
	})

	t.Run("defense", func(t *testing.T) {
		for _, tc := range []defenseTestcase{
			{
				title:   "こおり2倍",
				request: []byte(`{"ids":[6]}`),
				filter:  func(item item) bool { return floatEqual(item.Multiplier, 2) },
				want: items{
					newItem(2, "ほのお", "こおり"),
					newItem(2, "かくとう", "こおり"),
					newItem(2, "いわ", "こおり"),
					newItem(2, "はがね", "こおり"),
				},
			},
			{
				title:   "みずでんき2倍",
				request: []byte(`{"ids":[3,4]}`),
				filter:  func(item item) bool { return floatEqual(item.Multiplier, 2) },
				want: items{
					newItem(2, "くさ", "みず", "くさ", "でんき"),
					newItem(2, "じめん", "みず", "じめん", "でんき"),
				},
			},
		} {
			tc := tc
			t.Run(tc.title, func(t *testing.T) {
				tc.test(t, client)
			})
		}
	})
}

type (
	defenseRes = attackRes

	defenseTestcase struct {
		title   string
		request []byte
		filter  func(item) bool
		want    items
	}
)

func (tc *defenseTestcase) test(t *testing.T, c *client) {
	body, err := c.post("/defense", tc.request)
	if err != nil {
		t.Fatal(err)
	}
	var r defenseRes
	if err := json.Unmarshal(body, &r); err != nil {
		t.Fatal(err)
	}

	var got items
	for _, x := range r.Items {
		if tc.filter(x) {
			got = append(got, x)
		}
	}

	assert.Equal(t, tc.want.String(), got.String())
}

type (
	typeRes struct {
		Item struct {
			ID int `json:"id"`
		} `json:"item"`
	}

	nameCell struct {
		Name string `json:"name"`
	}

	pile struct {
		Attack  nameCell `json:"attack"`
		Defense nameCell `json:"defense"`
	}
	piles []pile

	item struct {
		Pile       piles   `json:"pile"`
		Multiplier float32 `json:"multiplier"`
	}
	items     []item
	attackRes struct {
		Items items `json:"items"`
	}
)

func newItem(multiplier float32, attackDefensePair ...string) item {
	piles := make([]pile, len(attackDefensePair)/2)
	var (
		i, j int
	)
	for i < len(attackDefensePair) {
		attack := attackDefensePair[i]
		i++
		defense := attackDefensePair[i]
		i++
		piles[j] = pile{
			Attack: nameCell{
				Name: attack,
			},
			Defense: nameCell{
				Name: defense,
			},
		}
		j++
	}
	return item{
		Multiplier: multiplier,
		Pile:       piles,
	}
}

func (p *pile) String() string {
	return fmt.Sprintf("%s-%s", p.Attack.Name, p.Defense.Name)
}
func (ps piles) String() string {
	ss := make([]string, len(ps))
	for i, x := range ps {
		ss[i] = x.String()
	}
	sort.Strings(ss)
	return strings.Join(ss, "|")
}
func (i *item) String() string {
	return fmt.Sprintf("%s_%f", i.Pile.String(), i.Multiplier)
}
func (is items) String() string {
	ss := make([]string, len(is))
	for i, x := range is {
		ss[i] = x.String()
	}
	sort.Strings(ss)
	return strings.Join(ss, "^")
}
func (a *attackRes) String() string {
	ss := make([]string, len(a.Items))
	for i, item := range a.Items {
		ss[i] = item.String()
	}
	sort.Strings(ss)
	return strings.Join(ss, ",")
}

type attackTestcase struct {
	title   string
	request []byte
	filter  func(item) bool
	want    items
}

func (tc *attackTestcase) test(t *testing.T, c *client) {
	body, err := c.post("/attack", tc.request)
	if err != nil {
		t.Fatal(err)
	}
	var r attackRes
	if err := json.Unmarshal(body, &r); err != nil {
		t.Fatal(err)
	}

	var got items
	for _, x := range r.Items {
		if tc.filter(x) {
			got = append(got, x)
		}
	}

	assert.Equal(t, tc.want.String(), got.String())
}

type typeTestcase struct {
	title   string
	request []byte
	wantID  int
}

func (tc *typeTestcase) test(t *testing.T, c *client) {
	body, err := c.post("/type", tc.request)
	if err != nil {
		t.Fatal(err)
	}
	var r typeRes
	if err := json.Unmarshal(body, &r); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, tc.wantID, r.Item.ID)
}

type client struct {
	client *http.Client
}

func (c *client) post(pattern string, body []byte) ([]byte, error) {
	r, err := c.client.Post(
		fmt.Sprintf("%s%s", endpoint, pattern),
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return io.ReadAll(r.Body)
}

func (c *client) init(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	if err := c.waitServerReady(t, ctx); err != nil {
		t.Fatalf("wait server ready: %v", err)
	}
}

func (c *client) waitServerReady(t *testing.T, ctx context.Context) error {
	t.Helper()
	request, _ := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if _, err := c.client.Do(request); err != nil {
				t.Log("Wait server ready...")
				time.Sleep(time.Second)
				continue
			}
			return nil
		}
	}
}

func runMain(pattern string) (func() error, error) {
	dir, err := os.MkdirTemp("", "server")
	if err != nil {
		return nil, err
	}

	executable := filepath.Join(dir, "server")
	if err := exec.Command("go", "build", "-o", executable, pattern).Run(); err != nil {
		os.RemoveAll(dir)
		return nil, err
	}

	cmd := exec.Command(executable)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return func() error {
		defer os.RemoveAll(dir)
		if err := cmd.Process.Signal(os.Interrupt); err != nil {
			return err
		}
		return cmd.Wait()
	}, nil
}
