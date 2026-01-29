Техническое задание для MVP ЗЕЛЕНИАЛ
1. User Service

Цель: регистрация, авторизация, управление профилем

Технологии:

Go (Golang) + Gin/Gorilla Mux для REST API

PostgreSQL для хранения данных

Redis для сессий

JWT + 2FA (TOTP)

Подсказки по реализации:

Структура проекта:

/cmd/userservice/main.go
/internal/app
/internal/models
/internal/repo
/internal/handler


Таблица пользователей PostgreSQL: id, email, password_hash, created_at, updated_at, 2fa_secret

Хэширование пароля: bcrypt

JWT для авторизации, refresh-токен для обновления сессии

API:

POST /register {email, password} → создаем пользователя
POST /login {email, password} → возвращаем JWT
GET /profile → возвращаем данные профиля
PUT /profile → редактируем профиль
POST /2fa/enable → генерируем TOTP
POST /2fa/verify → проверяем код


Добавить middleware для проверки токена на всех приватных эндпоинтах

2. Calendar & Tasks Service

Цель: управление задачами и расписанием

Технологии:

Go + Gin

PostgreSQL + Redis (кэш задач)

Kafka (очереди уведомлений)

Подсказки по реализации:

Таблицы:

tasks: id, user_id, title, description, due_date, status, created_at, updated_at

events: id, user_id, title, start_time, end_time, created_at, updated_at

API:

POST /tasks → создаем задачу
GET /tasks → список задач
PUT /tasks/{id} → редактируем задачу
DELETE /tasks/{id} → удаляем задачу
POST /events → создаем событие


Kafka: публикуем событие task_reminder перед дедлайном

Redis: хранение ближайших напоминаний для быстрого доступа

3. Finance Service

Цель: отслеживание доходов и расходов

Технологии:

Go + Gin

PostgreSQL

Simple Chart библиотека для веб

Подсказки по реализации:

Таблица transactions: id, user_id, type (income/expense), amount, category, date, description

API:

POST /transactions → добавить транзакцию
GET /transactions → список транзакций
GET /summary → сводка доходов/расходов


Прогноз бюджета: простой алгоритм на основе средней суммы расходов за прошлые 3 месяца

4. AI/Assistant Service

Цель: базовый чат и рекомендации

Технологии:

Go для API сервиса

LLM на сервере (можно использовать GPT-4, Llama или локальную LLM через REST API)

Kafka для получения событий из других сервисов

Шифрование данных пользователя

Подсказки по реализации:

API:

POST /assistant/chat → {user_id, message} → возвращает ответ AI


Обезличивание данных: все личные данные шифруются и не отправляются в AI, только обезличенные контексты

Примеры рекомендаций:

“Добавьте задачу в календарь на завтра 18:00 — полчаса отдыха”

“На этой неделе ваши расходы превысили план, сократите категории: кафе, развлечения”

5. Frontend MVP (веб)

Цель: интерфейс для взаимодействия с сервисами

Технологии:

React + TypeScript + Vite

Axios для API

TailwindCSS для стиля

Chart.js для графиков

Подсказки по реализации:

Структура:

/src
/components → кнопки, карточки задач, формы
/pages → Login, Register, Dashboard
/services → API запросы
/store → состояние пользователя, задач, транзакций


Dashboard:

Список задач и событий календаря

График расходов и доходов

Чат с AI ассистентом

Подключение к backend:

Axios с JWT токеном

Интерсептор для обработки ошибок

6. Инфраструктура и DevOps

Цель: готовность MVP к локальной разработке и облачному деплою

Технологии:

Docker для контейнеризации сервисов

Docker Compose для локальной среды

CI/CD: GitHub Actions для сборки, тестов, деплоя

Подсказки по реализации:

Dockerfile для каждого сервиса: сборка Go, установка зависимостей, запуск

docker-compose.yml:

services:
postgres:
image: postgres:15
redis:
image: redis:7
kafka:
image: confluentinc/cp-kafka:latest
user-service:
build: ./user-service
depends_on: [postgres, redis]


Настроить переменные окружения через .env

7. Тесты

Подсказки по реализации:

Go: testing + testify для юнит-тестов

Frontend: React Testing Library + Cypress

Проверка основных сценариев:

Регистрация и логин

Создание задачи

Добавление транзакции

Чат с AI

8. Документация

Подсказки:

API: Swagger/OpenAPI

README в каждом сервисе: как запускать, тестировать, деплоить

Архитектурная документация: блок-схемы, микросервисы, взаимодействие