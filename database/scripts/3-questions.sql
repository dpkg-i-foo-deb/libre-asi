--I12

INSERT INTO public.questions
(id, created_at, updated_at, deleted_at, question_category_id, "order", special_code, question_type_id, asi_form_id, dependency_id)
VALUES(1, '2023-04-30 16:30:46.267', NULL, NULL, 1, 1, 'I12', 1, 1, NULL);

INSERT INTO public.question_translations
(id, created_at, updated_at, deleted_at, "language", question_id, is_default, "statement", simplified_statement)
VALUES(1, '2023-04-30 16:31:27.385', NULL, NULL, 'ES', 1, NULL, '¿Cuál es su estado civil actual?', NULL);

INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(1, '2023-04-30 16:32:14.701', NULL, NULL, 1, 'ES', 1, 'Casado/a', 1, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(2, '2023-04-30 16:32:32.377', NULL, NULL, 1, 'ES', 2, 'Vive como si estuviera casado/a', 2, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(3, '2023-04-30 16:33:03.927', NULL, NULL, 1, 'ES', 3, 'Viudo/a', 3, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(4, '2023-04-30 16:33:22.330', NULL, NULL, 1, 'ES', 4, 'Divorciado/a', 4, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(5, '2023-04-30 16:33:38.612', NULL, NULL, 1, 'ES', 5, 'Separado/a', 5, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(6, '2023-04-30 16:33:55.083', NULL, NULL, 1, 'ES', 6, 'Nunca casado/a', 6, NULL);

--I13

INSERT INTO public.questions
(id, created_at, updated_at, deleted_at, question_category_id, "order", special_code, question_type_id, asi_form_id, dependency_id)
VALUES(2, '2023-04-30 16:34:56.459', NULL, NULL, 1, 2, 'I13', 2, 1, NULL);

INSERT INTO public.question_translations
(id, created_at, updated_at, deleted_at, "language", question_id, is_default, "statement", simplified_statement)
VALUES(3, '2023-04-30 16:41:07.955', NULL, NULL, NULL, 2, NULL, '¿Cuánto tiempo lleva en ese estado?', NULL);


INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(7, '2023-04-30 16:35:24.874', NULL, NULL, 2, 'ES', 1, 'Años', 99, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(8, '2023-04-30 16:38:27.679', NULL, NULL, 2, 'ES', 2, 'Meses', 99, NULL);


--I14

INSERT INTO public.questions
(id, created_at, updated_at, deleted_at, question_category_id, "order", special_code, question_type_id, asi_form_id, dependency_id)
VALUES(3, '2023-04-30 16:38:47.002', NULL, NULL, 1, 3, 'I14', 1, 1, NULL);

INSERT INTO public.question_translations
(id, created_at, updated_at, deleted_at, "language", question_id, is_default, "statement", simplified_statement)
VALUES(2, '2023-04-30 16:39:08.718', NULL, NULL, 'ES', 3, NULL, '¿Quién le derivó al tratamiento?', NULL);


INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(9, '2023-04-30 16:41:59.436', NULL, NULL, 3, 'ES', 1, 'Usted mismo, familia o amigos', 1, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(10, '2023-04-30 16:43:28.698', NULL, NULL, 3, 'ES', 2, 'Servicio o unidad de tratamiento de toxicomanías o alcoholismo', 2, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(11, '2023-04-30 16:44:07.864', NULL, NULL, 3, 'ES', 3, 'Médico de asistencia primaria', 3, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(12, '2023-04-30 16:44:30.390', NULL, NULL, 3, 'ES', 4, 'Otros servicios sanitarios', 4, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(13, '2023-04-30 16:44:51.255', NULL, NULL, 3, 'ES', 5, 'Sistema educativo', 5, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(14, '2023-04-30 16:45:09.531', NULL, NULL, 3, 'ES', 6, 'Mutua laboral', 6, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(15, '2023-04-30 16:45:28.430', NULL, NULL, 3, 'ES', 7, 'Agentes sociales (INEM, centro de acogida, iglesia, etc)', 7, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, value, simplified_description)
VALUES(16, '2023-04-30 16:45:54.848', NULL, NULL, 3, 'ES', 8, 'Sistema legal o judicial', 8, NULL);

-- I15

