Run these from the same directory where your `docker-compose.yml` file lives:

---

## 1. Start everything up

```bash
docker compose up -d

```

* **What it does:** Downloads the Postgres image (if you don't have it), creates the container, and starts the database in the background.
* **The `-d` flag:** Stands for "detached" mode. It keeps the database running quietly in the background so your terminal remains free for you to type other commands.

## 2. Check if it's running

```bash
docker compose ps

```

* **What it does:** Lists the containers managed by this specific YAML file.
* **What to look for:** You want to see your container status as `Up` or `running` and verify that port `5432` is exposed.

## 3. View the database logs

```bash
docker compose logs -f postgres

```

* **What it does:** Shows you what is happening inside Postgres. If your Go app fails to connect or a query crashes the database, the error messages will print here.
* **The `-f` flag:** Stands for "follow". It streams live updates as they happen. Press `Ctrl + C` to exit the log view (this won't stop the database).

## 4. Stop everything (without losing data)

```bash
docker compose down

```

* **What it does:** Safely stops and removes the container. Because we set up a `volume` in your YAML file, your tables and data will still be waiting for you the next time you run `docker compose up`.

## 5. Nuclear Option: Wipe everything and start fresh

```bash
docker compose down -v

```

* **What it does:** The `-v` flag deletes the volumes along with the container.
* **When to use it:** If you mess up your database design halfway through Day 1, panic, and want to completely wipe all your tables and data to start over with a blank canvas, run this.

---