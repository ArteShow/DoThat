document.addEventListener("DOMContentLoaded", () => {
    const path = window.location.pathname;

    if (path === "/" || path === "/index.html") {
        loadDashboard();
    }

    if (path === "/tasks.html") {
        loadTasksPage();
    }

    const refreshButton = document.querySelector("#refresh-tasks");

    if (refreshButton) {
        refreshButton.addEventListener("click", loadTasksPage);
    }
});

async function loadDashboard() {
    const container = document.querySelector("#upcoming-tasks");

    try {
        const data = await getTasks();
        const tasks = data.tasks ?? [];

        document.querySelector("#upcoming-count").textContent =
            tasks.filter(task => task.status !== "completed").length;

        document.querySelector("#completed-count").textContent =
            tasks.filter(task => task.status === "completed").length;

        document.querySelector("#schedule-count").textContent = "—";

        renderTasks(tasks, container, 5);
    } catch (error) {
        container.textContent = "Couldn't load tasks. Check that the DoThat! server is running.";
        console.error(error);
    }
}

async function loadTasksPage() {
    const container = document.querySelector("#all-tasks");

    if (!container) return;

    container.textContent = "Loading tasks...";

    try {
        const data = await getTasks();
        renderTasks(data.tasks ?? [], container);
    } catch (error) {
        container.textContent = "Couldn't load tasks. Check the API connection.";
        console.error(error);
    }
}

function renderTasks(tasks, container, limit = Infinity) {
    container.replaceChildren();

    if (tasks.length === 0) {
        container.textContent = "No tasks yet. Enjoy the free time!";
        container.classList.add("muted");
        return;
    }

    container.classList.remove("muted");

    tasks.slice(0, limit).forEach(task => {
        const row = document.createElement("div");
        row.className = "task-row";

        const details = document.createElement("div");

        const title = document.createElement("div");
        title.className = "task-title";
        title.textContent = task.title ?? "Untitled task";

        const meta = document.createElement("div");
        meta.className = "task-meta";
        meta.textContent = task.deadline
            ? `Deadline: ${task.deadline}`
            : "No deadline";

        details.append(title, meta);

        const status = document.createElement("span");
        status.className = "task-status";
        status.textContent = task.status ?? "pending";

        row.append(details, status);
        container.append(row);
    });
}
