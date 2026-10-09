document.addEventListener("DOMContentLoaded", () => {
    const path = window.location.pathname;

    if (path === "/" || path === "/index.html") {
        loadDashboard();
    }

    if (path === "/tasks.html") {
        loadTasksPage();

        document
            .querySelector("#create-task-form")
            .addEventListener("submit", handleCreateTask);

        document
            .querySelector("#refresh-tasks")
            .addEventListener("click", loadTasksPage);
    }
});

function getTaskStatus(task) {
    return String(task.status ?? "pending").toLowerCase();
}

function isCompleted(task) {
    return getTaskStatus(task) === "completed";
}

function isOverdue(task) {
    if (!task.deadline || isCompleted(task)) {
        return false;
    }

    const deadline = new Date(task.deadline);

    return !Number.isNaN(deadline.getTime()) && deadline < new Date();
}

function formatDeadline(deadline) {
    if (!deadline) {
        return "No deadline";
    }

    const date = new Date(deadline);

    if (Number.isNaN(date.getTime())) {
        return `Deadline: ${deadline}`;
    }

    return `Due ${date.toLocaleString([], {
        dateStyle: "medium",
        timeStyle: "short",
    })}`;
}

async function loadDashboard() {
    const container = document.querySelector("#upcoming-tasks");

    try {
        const data = await apiGetTasks();
        const tasks = data?.tasks ?? [];

        document.querySelector("#upcoming-count").textContent =
            tasks.filter(task => !isCompleted(task)).length;

        document.querySelector("#completed-count").textContent =
            tasks.filter(isCompleted).length;

        document.querySelector("#overdue-count").textContent =
            tasks.filter(isOverdue).length;

        const sortedTasks = [...tasks].sort((a, b) => {
            if (!a.deadline) return 1;
            if (!b.deadline) return -1;

            return new Date(a.deadline) - new Date(b.deadline);
        });

        const upcoming = sortedTasks.filter(task => !isCompleted(task));

        renderTasks(upcoming, container, 5);
    } catch (error) {
        showError(
            container,
            "Couldn't load tasks. Check that the Go server and API are running."
        );
        console.error("Dashboard:", error);
    }
}

async function loadTasksPage() {
    const container = document.querySelector("#all-tasks");
    const message = document.querySelector("#tasks-message");

    if (!container) return;

    container.textContent = "Loading tasks...";
    clearMessage(message);

    try {
        const data = await apiGetTasks();
        const tasks = data?.tasks ?? [];

        const sortedTasks = [...tasks].sort((a, b) => {
            if (isCompleted(a) !== isCompleted(b)) {
                return isCompleted(a) ? 1 : -1;
            }

            if (!a.deadline) return 1;
            if (!b.deadline) return -1;

            return new Date(a.deadline) - new Date(b.deadline);
        });

        renderTasks(sortedTasks, container);
    } catch (error) {
        showError(container, "Couldn't load tasks. Check the API connection.");
        console.error("Tasks:", error);
    }
}

async function handleCreateTask(event) {
    event.preventDefault();

    const form = event.currentTarget;
    const button = form.querySelector('button[type="submit"]');
    const message = document.querySelector("#create-task-message");
    const formData = new FormData(form);

    const title = String(formData.get("title") ?? "").trim();

    if (!title) {
        showMessage(message, "Please enter a task title.", "error");
        return;
    }

    const rawDeadline = String(formData.get("deadline") ?? "");
    const rawDuration = String(formData.get("duration") ?? "");

    let deadline = null;

    if (rawDeadline) {
        const parsedDeadline = new Date(rawDeadline);

        if (Number.isNaN(parsedDeadline.getTime())) {
            showMessage(message, "Please enter a valid deadline.", "error");
            return;
        }

        // Send an RFC3339-compatible timestamp to Go.
        deadline = parsedDeadline.toISOString();
    }

    const task = {
        title,
        description: String(formData.get("description") ?? "").trim(),
        deadline,
        duration: rawDuration ? Number(rawDuration) : 0,
        priority: Number(formData.get("priority") ?? 0),
        status: "pending",
    };

    button.disabled = true;
    showMessage(message, "Creating task...");

    try {
        await apiCreateTask(task);

        form.reset();
        showMessage(message, "Task created successfully.", "success");

        await loadTasksPage();
    } catch (error) {
        showMessage(message, `Couldn't create task: ${error.message}`, "error");
        console.error("Create task:", error);
    } finally {
        button.disabled = false;
    }
}

function renderTasks(tasks, container, limit = Infinity) {
    container.replaceChildren();

    if (tasks.length === 0) {
        const empty = document.createElement("p");
        empty.className = "muted";
        empty.textContent = "No tasks here yet.";
        container.append(empty);
        return;
    }

    for (const task of tasks.slice(0, limit)) {
        const row = document.createElement("article");
        row.className = "task-row";

        if (isCompleted(task)) {
            row.classList.add("is-completed");
        }

        const details = document.createElement("div");
        details.className = "task-details";

        const title = document.createElement("div");
        title.className = "task-title";
        title.textContent = task.title ?? "Untitled task";

        const description = document.createElement("p");
        description.className = "task-description";
        description.textContent = task.description ?? "";

        if (!description.textContent) {
            description.hidden = true;
        }

        const meta = document.createElement("div");
        meta.className = "task-meta";

        const deadlineText = formatDeadline(task.deadline);
        const priorityText = `Priority: ${task.priority ?? 0}`;
        const durationText = task.duration
            ? `${task.duration} min`
            : "";

        meta.textContent = [deadlineText, priorityText, durationText]
            .filter(Boolean)
            .join(" · ");

        details.append(title, description, meta);

        const actions = document.createElement("div");
        actions.className = "task-actions";

        const status = document.createElement("span");
        status.className = "task-status";
        status.textContent = getTaskStatus(task);
        actions.append(status);

        if (task.id) {
            const toggleButton = document.createElement("button");
            toggleButton.type = "button";
            toggleButton.className = "small-button";
            toggleButton.textContent = isCompleted(task)
                ? "Reopen"
                : "Complete";

            toggleButton.addEventListener("click", async () => {
                await changeTaskStatus(
                    task.id,
                    isCompleted(task) ? "pending" : "completed"
                );
            });

            actions.append(toggleButton);

            const deleteButton = document.createElement("button");
            deleteButton.type = "button";
            deleteButton.className = "small-button delete-button";
            deleteButton.textContent = "Delete";

            deleteButton.addEventListener("click", async () => {
                if (!window.confirm(`Delete "${task.title ?? "this task"}"?`)) {
                    return;
                }

                await removeTask(task.id);
            });

            actions.append(deleteButton);
        }

        row.append(details, actions);
        container.append(row);
    }
}

async function changeTaskStatus(taskID, newStatus) {
    const message = document.querySelector("#tasks-message");

    try {
        await apiUpdateTaskStatus(taskID, newStatus);
        showMessage(message, "Task status updated.", "success");

        if (window.location.pathname === "/") {
            await loadDashboard();
        } else {
            await loadTasksPage();
        }
    } catch (error) {
        showMessage(message, `Couldn't update task: ${error.message}`, "error");
        console.error("Update status:", error);
    }
}

async function removeTask(taskID) {
    const message = document.querySelector("#tasks-message");

    try {
        await apiDeleteTask(taskID);
        showMessage(message, "Task deleted.", "success");

        if (window.location.pathname === "/") {
            await loadDashboard();
        } else {
            await loadTasksPage();
        }
    } catch (error) {
        showMessage(message, `Couldn't delete task: ${error.message}`, "error");
        console.error("Delete task:", error);
    }
}

function showError(container, text) {
    container.replaceChildren();

    const paragraph = document.createElement("p");
    paragraph.className = "form-message error";
    paragraph.textContent = text;

    container.append(paragraph);
}

function showMessage(element, text, type = "") {
    if (!element) return;

    element.textContent = text;
    element.className = `form-message ${type}`.trim();
}

function clearMessage(element) {
    if (!element) return;

    element.textContent = "";
    element.className = "form-message";
}