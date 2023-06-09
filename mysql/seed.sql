drop table if exists effectivities;
drop table if exists types;

create table if not exists types (
  id int primary key,
  name varchar(255) not null
);

insert into types (id, name) values
(1, 'ノーマル'),
(2, 'ほのお'),
(3, 'みず'),
(4, 'でんき'),
(5, 'くさ'),
(6, 'こおり'),
(7, 'かくとう'),
(8, 'どく'),
(9, 'じめん'),
(10, 'ひこう'),
(11, 'エスパー'),
(12, 'むし'),
(13, 'いわ'),
(14, 'ゴースト'),
(15, 'ドラゴン'),
(16, 'あく'),
(17, 'はがね'),
(18, 'フェアリー')
;

create table if not exists effectivities (
  id int primary key,
  attack_type_id int,
  defense_type_id int,
  multiplier double not null,
  foreign key fk_attack_type_id (attack_type_id) references types (id),
  foreign key fk_defense_type_id (defense_type_id) references types (id)
);

insert into effectivities (id, attack_type_id, defense_type_id, multiplier) values
(1, 1, 1, 1),
(2, 1, 2, 1),
(3, 1, 3, 1),
(4, 1, 4, 1),
(5, 1, 5, 1),
(6, 1, 6, 1),
(7, 1, 7, 1),
(8, 1, 8, 1),
(9, 1, 9, 1),
(10, 1, 10, 1),
(11, 1, 11, 1),
(12, 1, 12, 1),
(13, 1, 13, 0.5),
(14, 1, 14, 0),
(15, 1, 15, 1),
(16, 1, 16, 1),
(17, 1, 17, 0.5),
(18, 1, 18, 1),
(19, 2, 1, 1),
(20, 2, 2, 0.5),
(21, 2, 3, 0.5),
(22, 2, 4, 1),
(23, 2, 5, 2),
(24, 2, 6, 2),
(25, 2, 7, 1),
(26, 2, 8, 1),
(27, 2, 9, 1),
(28, 2, 10, 1),
(29, 2, 11, 1),
(30, 2, 12, 2),
(31, 2, 13, 0.5),
(32, 2, 14, 1),
(33, 2, 15, 0.5),
(34, 2, 16, 1),
(35, 2, 17, 2),
(36, 2, 18, 1),
(37, 3, 1, 1),
(38, 3, 2, 2),
(39, 3, 3, 0.5),
(40, 3, 4, 1),
(41, 3, 5, 0.5),
(42, 3, 6, 1),
(43, 3, 7, 1),
(44, 3, 8, 1),
(45, 3, 9, 2),
(46, 3, 10, 1),
(47, 3, 11, 1),
(48, 3, 12, 1),
(49, 3, 13, 2),
(50, 3, 14, 1),
(51, 3, 15, 0.5),
(52, 3, 16, 1),
(53, 3, 17, 1),
(54, 3, 18, 1),
(55, 4, 1, 1),
(56, 4, 2, 1),
(57, 4, 3, 2),
(58, 4, 4, 0.5),
(59, 4, 5, 0.5),
(60, 4, 6, 1),
(61, 4, 7, 1),
(62, 4, 8, 1),
(63, 4, 9, 0),
(64, 4, 10, 2),
(65, 4, 11, 1),
(66, 4, 12, 1),
(67, 4, 13, 1),
(68, 4, 14, 1),
(69, 4, 15, 0.5),
(70, 4, 16, 1),
(71, 4, 17, 1),
(72, 4, 18, 1),
(73, 5, 1, 1),
(74, 5, 2, 0.5),
(75, 5, 3, 2),
(76, 5, 4, 1),
(77, 5, 5, 0.5),
(78, 5, 6, 1),
(79, 5, 7, 1),
(80, 5, 8, 0.5),
(81, 5, 9, 2),
(82, 5, 10, 0.5),
(83, 5, 11, 1),
(84, 5, 12, 0.5),
(85, 5, 13, 2),
(86, 5, 14, 1),
(87, 5, 15, 0.5),
(88, 5, 16, 1),
(89, 5, 17, 0.5),
(90, 5, 18, 1),
(91, 6, 1, 1),
(92, 6, 2, 0.5),
(93, 6, 3, 0.5),
(94, 6, 4, 1),
(95, 6, 5, 2),
(96, 6, 6, 0.5),
(97, 6, 7, 1),
(98, 6, 8, 1),
(99, 6, 9, 2),
(100, 6, 10, 2),
(101, 6, 11, 1),
(102, 6, 12, 1),
(103, 6, 13, 1),
(104, 6, 14, 1),
(105, 6, 15, 2),
(106, 6, 16, 1),
(107, 6, 17, 0.5),
(108, 6, 18, 1),
(109, 7, 1, 2),
(110, 7, 2, 1),
(111, 7, 3, 1),
(112, 7, 4, 1),
(113, 7, 5, 1),
(114, 7, 6, 2),
(115, 7, 7, 1),
(116, 7, 8, 0.5),
(117, 7, 9, 1),
(118, 7, 10, 0.5),
(119, 7, 11, 0.5),
(120, 7, 12, 0.5),
(121, 7, 13, 2),
(122, 7, 14, 0),
(123, 7, 15, 1),
(124, 7, 16, 2),
(125, 7, 17, 2),
(126, 7, 18, 0.5),
(127, 8, 1, 1),
(128, 8, 2, 1),
(129, 8, 3, 1),
(130, 8, 4, 1),
(131, 8, 5, 2),
(132, 8, 6, 1),
(133, 8, 7, 1),
(134, 8, 8, 0.5),
(135, 8, 9, 0.5),
(136, 8, 10, 1),
(137, 8, 11, 1),
(138, 8, 12, 1),
(139, 8, 13, 0.5),
(140, 8, 14, 0.5),
(141, 8, 15, 1),
(142, 8, 16, 1),
(143, 8, 17, 0),
(144, 8, 18, 2),
(145, 9, 1, 1),
(146, 9, 2, 2),
(147, 9, 3, 1),
(148, 9, 4, 2),
(149, 9, 5, 0.5),
(150, 9, 6, 1),
(151, 9, 7, 1),
(152, 9, 8, 2),
(153, 9, 9, 1),
(154, 9, 10, 0),
(155, 9, 11, 1),
(156, 9, 12, 0.5),
(157, 9, 13, 2),
(158, 9, 14, 1),
(159, 9, 15, 1),
(160, 9, 16, 1),
(161, 9, 17, 2),
(162, 9, 18, 1),
(163, 10, 1, 1),
(164, 10, 2, 1),
(165, 10, 3, 1),
(166, 10, 4, 0.5),
(167, 10, 5, 2),
(168, 10, 6, 1),
(169, 10, 7, 2),
(170, 10, 8, 1),
(171, 10, 9, 1),
(172, 10, 10, 1),
(173, 10, 11, 1),
(174, 10, 12, 2),
(175, 10, 13, 0.5),
(176, 10, 14, 1),
(177, 10, 15, 1),
(178, 10, 16, 1),
(179, 10, 17, 0.5),
(180, 10, 18, 1),
(181, 11, 1, 1),
(182, 11, 2, 1),
(183, 11, 3, 1),
(184, 11, 4, 1),
(185, 11, 5, 1),
(186, 11, 6, 1),
(187, 11, 7, 2),
(188, 11, 8, 2),
(189, 11, 9, 1),
(190, 11, 10, 1),
(191, 11, 11, 0.5),
(192, 11, 12, 1),
(193, 11, 13, 1),
(194, 11, 14, 1),
(195, 11, 15, 1),
(196, 11, 16, 0),
(197, 11, 17, 0.5),
(198, 11, 18, 1),
(199, 12, 1, 1),
(200, 12, 2, 0.5),
(201, 12, 3, 1),
(202, 12, 4, 1),
(203, 12, 5, 2),
(204, 12, 6, 1),
(205, 12, 7, 0.5),
(206, 12, 8, 0.5),
(207, 12, 9, 1),
(208, 12, 10, 0.5),
(209, 12, 11, 2),
(210, 12, 12, 1),
(211, 12, 13, 1),
(212, 12, 14, 0.5),
(213, 12, 15, 1),
(214, 12, 16, 2),
(215, 12, 17, 0.5),
(216, 12, 18, 0.5),
(217, 13, 1, 1),
(218, 13, 2, 2),
(219, 13, 3, 1),
(220, 13, 4, 1),
(221, 13, 5, 1),
(222, 13, 6, 2),
(223, 13, 7, 0.5),
(224, 13, 8, 1),
(225, 13, 9, 0.5),
(226, 13, 10, 2),
(227, 13, 11, 1),
(228, 13, 12, 2),
(229, 13, 13, 1),
(230, 13, 14, 1),
(231, 13, 15, 1),
(232, 13, 16, 1),
(233, 13, 17, 0.5),
(234, 13, 18, 1),
(235, 14, 1, 0),
(236, 14, 2, 1),
(237, 14, 3, 1),
(238, 14, 4, 1),
(239, 14, 5, 1),
(240, 14, 6, 1),
(241, 14, 7, 1),
(242, 14, 8, 1),
(243, 14, 9, 1),
(244, 14, 10, 1),
(245, 14, 11, 2),
(246, 14, 12, 1),
(247, 14, 13, 1),
(248, 14, 14, 2),
(249, 14, 15, 1),
(250, 14, 16, 0.5),
(251, 14, 17, 1),
(252, 14, 18, 1),
(253, 15, 1, 1),
(254, 15, 2, 1),
(255, 15, 3, 1),
(256, 15, 4, 1),
(257, 15, 5, 1),
(258, 15, 6, 1),
(259, 15, 7, 1),
(260, 15, 8, 1),
(261, 15, 9, 1),
(262, 15, 10, 1),
(263, 15, 11, 1),
(264, 15, 12, 1),
(265, 15, 13, 1),
(266, 15, 14, 1),
(267, 15, 15, 2),
(268, 15, 16, 1),
(269, 15, 17, 0.5),
(270, 15, 18, 0),
(271, 16, 1, 1),
(272, 16, 2, 1),
(273, 16, 3, 1),
(274, 16, 4, 1),
(275, 16, 5, 1),
(276, 16, 6, 1),
(277, 16, 7, 0.5),
(278, 16, 8, 1),
(279, 16, 9, 1),
(280, 16, 10, 1),
(281, 16, 11, 2),
(282, 16, 12, 1),
(283, 16, 13, 1),
(284, 16, 14, 2),
(285, 16, 15, 1),
(286, 16, 16, 0.5),
(287, 16, 17, 1),
(288, 16, 18, 0.5),
(289, 17, 1, 1),
(290, 17, 2, 0.5),
(291, 17, 3, 0.5),
(292, 17, 4, 0.5),
(293, 17, 5, 1),
(294, 17, 6, 2),
(295, 17, 7, 1),
(296, 17, 8, 1),
(297, 17, 9, 1),
(298, 17, 10, 1),
(299, 17, 11, 1),
(300, 17, 12, 1),
(301, 17, 13, 2),
(302, 17, 14, 1),
(303, 17, 15, 1),
(304, 17, 16, 1),
(305, 17, 17, 0.5),
(306, 17, 18, 2),
(307, 18, 1, 1),
(308, 18, 2, 0.5),
(309, 18, 3, 1),
(310, 18, 4, 1),
(311, 18, 5, 1),
(312, 18, 6, 1),
(313, 18, 7, 2),
(314, 18, 8, 0.5),
(315, 18, 9, 1),
(316, 18, 10, 1),
(317, 18, 11, 1),
(318, 18, 12, 1),
(319, 18, 13, 1),
(320, 18, 14, 1),
(321, 18, 15, 2),
(322, 18, 16, 2),
(323, 18, 17, 0.5),
(324, 18, 18, 1)
;
