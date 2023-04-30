-- Auto-generated SQL script #202304301602
INSERT INTO public.questions (created_at,question_category_id,"order",special_code,question_type_id,asi_form_id)
	VALUES (NOW(),1,1,'I12',9,1);

INSERT INTO public.question_translations (created_at,"language",question_id,"statement")
	VALUES ('2023-04-30 16:04:03.753','ES',1,'¿Cuál es su estado civil actual?');

INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, simplified_description, dependency_id)
VALUES(1, '2023-04-30 16:06:59.895', NULL, NULL, 1, 'ES', 1, 'Casado/a', NULL, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, simplified_description, dependency_id)
VALUES(2, '2023-04-30 16:07:33.219', NULL, NULL, 1, 'ES', 2, 'Vive como si estuviera casado/a', NULL, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, simplified_description, dependency_id)
VALUES(3, '2023-04-30 16:07:56.913', NULL, NULL, 1, 'ES', 4, 'Divorciado/a', NULL, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, simplified_description, dependency_id)
VALUES(4, '2023-04-30 16:08:20.690', NULL, NULL, 1, 'ES', 5, 'Separado/a', NULL, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, simplified_description, dependency_id)
VALUES(5, '2023-04-30 16:08:43.132', NULL, NULL, 1, 'ES', 6, 'Nunca casado/a', NULL, NULL);
INSERT INTO public."options"
(id, created_at, updated_at, deleted_at, question_id, "language", "order", description, simplified_description, dependency_id)
VALUES(6, '2023-04-30 16:09:07.921', NULL, NULL, 1, 'ES', 3, 'Viudo/a', NULL, NULL);


INSERT INTO public.questions
(id, created_at, updated_at, deleted_at, question_category_id, "order", special_code, question_type_id, asi_form_id)
VALUES(2, '2023-04-30 16:05:27.635', NULL, NULL, 1, 2, 'I13', 10, 1);
