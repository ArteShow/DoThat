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
        throw new Error(`API request failed: ${response.status}`);
    }

    if (response.status === 204) {
        return null;
    }

    return response.json();
}

async function getTasks() {
    return apiRequest("/tasks/");
}
