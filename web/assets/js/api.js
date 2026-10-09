const API_BASE = "/api/v1";

async function apiRequest(path, options = {}) {
    const response = await fetch(`${API_BASE}${path}`, {
        ...options,
        headers: {
            "Content-Type": "application/json",
            ...options.headers,
        },
    });

    if (!response.ok) {
        const message = await response.text();
        throw new Error(message || `Request failed: ${response.status}`);
    }

    if (response.status === 204) {
        return null;
    }

    const text = await response.text();
    return text ? JSON.parse(text) : null;
}

function apiGetTasks() {
    return apiRequest("/tasks/");
}

function apiCreateTask(task) {
    return apiRequest("/tasks/", {
        method: "POST",
        body: JSON.stringify(task),
    });
}

function apiDeleteTask(taskID) {
    return apiRequest("/tasks/", {
        method: "DELETE",
        body: JSON.stringify({
            task_id: taskID,
        }),
    });
}

function apiUpdateTaskStatus(taskID, newStatus) {
    return apiRequest("/tasks/status", {
        method: "PATCH",
        body: JSON.stringify({
            task_id: taskID,
            new_status: newStatus,
        }),
    });
}

function apiGetPlannerEntries() {
    return apiRequest("/planner/");
}

function apiCreatePlannerEntry(entry) {
    return apiRequest("/planner/", {
        method: "POST",
        body: JSON.stringify(entry),
    });
}

function apiDeletePlannerEntry(entryID) {
    return apiRequest("/planner/", {
        method: "DELETE",
        body: JSON.stringify({
            entry_id: entryID,
        }),
    });
}
