import psycopg2
from faker import Faker
import random
import datetime

DBNAME = 'rk1-1'
USER = 'ivan'
PASSWORD = 'ivan2003'
HOST = 'localhost'
PORT = 5432

ELEMENTS = 10001

fake = Faker()
conn = psycopg2.connect(f"dbname={DBNAME} user={USER} password={PASSWORD} host={HOST} port={PORT}")

cur = conn.cursor()

# Заполнение таблицы users
for i in range(1, ELEMENTS):
    email = fake.email()
    password_hash = fake.password()
    avatar_url = fake.image_url()
    sub_expiration = datetime.datetime.now() + datetime.timedelta(days=30)
    created_at = fake.date_time_this_year()
    updated_at = created_at + datetime.timedelta(days=random.randint(1, 365), seconds=random.randint(1, 86400))
    cur.execute("INSERT INTO user_schema.users (email, password_hash, avatar_url, sub_expiration, created_at, updated_at) VALUES (%s, %s, %s, %s, %s, %s) ON CONFLICT DO NOTHING", (email, password_hash, avatar_url, sub_expiration, created_at, updated_at))

# Заполнение таблиц с контентом
countries = ['USA', 'UK', 'Germany', 'France', 'Spain', 'Italy', 'Russia', 'China', 'Japan', 'South Korea']
genres = ['Action', 'Comedy', 'Drama', 'Horror', 'Romance', 'Science Fiction', 'Thriller']
roles = ['Actor', 'Director', 'Writer', 'Producer']

for i in range(1, ELEMENTS):
    title = fake.sentence(nb_words=5)
    description = fake.paragraph(nb_sentences=3)
    rating = round(random.uniform(1, 10), 1)
    year = random.randint(1950, 2021)
    is_free = random.choice([True, False])
    age_limit = random.randint(0, 18)
    preview_url = fake.image_url()
    trailer_url = fake.url()
    type = random.choice(['film', 'series'])
    sum_ratings = random.randint(100, 10000)
    count_ratings = random.randint(10, 1000)

    cur.execute("INSERT INTO content_schema.content (title, description, rating, year, is_free, age_limit, preview_url, trailer_url, type, sum_ratings, count_ratings) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", (title, description, rating, year, is_free, age_limit, preview_url, trailer_url, type, sum_ratings, count_ratings))
    conn.commit()

    content_id = cur.lastrowid + 1

    # Добавление жанров и стран, связанных с контентом
    for j in range(random.randint(1, 5)):
        genre = random.choice(genres)
        country = random.choice(countries)

        cur.execute("INSERT INTO content_schema.genres (name) VALUES (%s) ON CONFLICT DO NOTHING", (genre,))
        cur.execute("INSERT INTO content_schema.countries (name) VALUES (%s) ON CONFLICT DO NOTHING", (country,))

        genre_id_query = "SELECT id FROM content_schema.genres WHERE name = %s"
        country_id_query = "SELECT id FROM content_schema.countries WHERE name = %s"
        cur.execute(genre_id_query, (genre,))
        genre_id = cur.fetchone()[0]

        cur.execute(country_id_query, (country,))
        country_id = cur.fetchone()[0]

        cur.execute("INSERT INTO content_schema.content_genres (content_id, genre_id) VALUES (%s, %s) ON CONFLICT DO NOTHING", (content_id, genre_id))
        cur.execute("INSERT INTO content_schema.content_countries (content_id, country_id) VALUES (%s, %s) ON CONFLICT DO NOTHING", (content_id, country_id))

    # Добавление ролей и людей, связанных с контентом
    for j in range(random.randint(1, 15)):
        role = random.choice(roles)
        name = fake.name()
        gender = random.choice(['M', 'F'])
        growth = random.randint(150, 200)
        birthplace = random.choice(countries)
        avatar_url = fake.image_url()
        age = random.randint(18, 80)

        cur.execute("INSERT INTO content_schema.roles (title) VALUES (%s) ON CONFLICT DO NOTHING", (role,))
        cur.execute("INSERT INTO content_schema.persons (name, gender, growth, birthplace, avatar_url, age) VALUES (%s, %s, %s, %s, %s, %s) RETURNING id", (name, gender, growth, birthplace, avatar_url, age))
        person_id = cur.fetchone()[0]

        role_id_query = "SELECT id FROM content_schema.roles WHERE title = %s"
        cur.execute(role_id_query, (role,))
        role_id = cur.fetchone()[0]

        cur.execute("INSERT INTO content_schema.content_roles_persons (role_id, person_id, content_id) VALUES (%s, %s, %s) ON CONFLICT DO NOTHING", (role_id, person_id, content_id))
    conn.commit()

    # Добавление эпизодов и сезонов, связанных с контентом, если это телесериал
    if type == 'series':
        series_id_query = "INSERT INTO content_schema.series (content_id) VALUES (%s) RETURNING id"
        cur.execute(series_id_query, (content_id,))
        series_id = cur.fetchone()[0]

        for j in range(random.randint(1, 6)):
            season_num = j + 1

            for k in range(random.randint(5, 20)):
                title = fake.sentence(nb_words=5)
                preview_url = fake.image_url()
                episode_num = k + 1
                content_url = fake.url()

                cur.execute("INSERT INTO content_schema.episodes (series_id, preview_url, season_num, episode_num, content_url, title) VALUES (%s, %s, %s, %s, %s, %s)", (series_id, preview_url, season_num, episode_num, content_url, title))

  # Добавление избранных контентов пользователей
for i in range(1, ELEMENTS):
    for j in range(random.randint(5, 20)):
        user_id = i
        content_id = random.randint(1, 100)

        cur.execute("INSERT INTO user_action_schema.users_content_favorites (user_id, content_id) VALUES (%s, %s) ON CONFLICT DO NOTHING", (user_id, content_id))

# Добавление избранных персон пользователей
for i in range(1, ELEMENTS):
    for j in range(random.randint(5, 20)):
        user_id = i
        person_id = random.randint(1, 100)

        cur.execute("INSERT INTO user_action_schema.users_persons_favorites (user_id, person_id) VALUES (%s, %s) ON CONFLICT DO NOTHING", (user_id, person_id))

# Добавление истории просмотра пользователей
for i in range(1, ELEMENTS):
    content_ids = random.sample(range(1, 101), random.randint(10, 50))
    stop_view = 0
    duration = 0

    for content_id in content_ids:
        duration = random.randint(10, 600)
        stop_view += duration
        cur.execute("INSERT INTO user_action_schema.history_views (user_id, content_id, stop_view, duration) VALUES (%s, %s, %s, %s)", (i, content_id, stop_view, duration))

# Добавление рейтинга контента пользователями
for i in range(1, ELEMENTS):
    content_ids = random.sample(range(1, 101), random.randint(10, 50))
    ratings = []

    for content_id in content_ids:
        rating = round(random.uniform(1, 10), 1)
        ratings.append((i, content_id, rating))

    cur.executemany("INSERT INTO user_action_schema.ratings (user_id, content_id, rating) VALUES (%s, %s, %s) ON CONFLICT DO NOTHING", ratings)
    conn.commit()

# Добавление сборников контента
selections = ['New Releases', 'Popular Now', 'Recommended for You', 'Staff Picks']

for selection in selections:
    cur.execute("INSERT INTO content_schema.selections (title) VALUES (%s) ON CONFLICT DO NOTHING", (selection,))
    conn.commit()

selection_ids_query = "SELECT id FROM content_schema.selections"

cur.execute(selection_ids_query)
selection_ids = cur.fetchall()

for content_id in range(1, ELEMENTS):
    selection_id = random.choice(selection_ids)[0]
    cur.execute("INSERT INTO content_schema.content_selections(content_id, selection_id) VALUES (%s, %s) ON CONFLICT DO NOTHING", (content_id, selection_id))

conn.commit()
# Закрытие соединения с базой данных
cur.close()
conn.close()