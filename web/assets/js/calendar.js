document.addEventListener("DOMContentLoaded", () => {
    const state = {
        mode: "timetable",
        date: new Date(),
        entries: [],
        tasks: [],
        selectedDay: null
    };

    const view = document.querySelector("#calendar-view");
    const heading = document.querySelector("#calendar-heading");
    const status = document.querySelector("#calendar-status");
    const formCard = document.querySelector("#create-event-card");
    const form = document.querySelector("#create-event-form");

    function setMessage(message, type = "") {
        status.textContent = message;
        status.className = `form-message ${type}`.trim();
    }

    function formatDate(date, options = {}) {
        return date.toLocaleDateString(undefined, options);
    }

    function startOfWeek(date) {
        const result = new Date(date);
        result.setHours(0, 0, 0, 0);

        const day = (result.getDay() + 6) % 7;
        result.setDate(result.getDate() - day);

        return result;
    }

    function dateKey(date) {
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, "0");
        const day = String(date.getDate()).padStart(2, "0");

        return `${year}-${month}-${day}`;
    }

    function parseEntryDate(value) {
        if (!value) return null;

        const date = new Date(value);
        return Number.isNaN(date.getTime()) ? null : date;
    }

    function getEntryStart(entry) {
        return parseEntryDate(entry.start_time);
    }

    function getEntryEnd(entry) {
        return parseEntryDate(entry.end_time);
    }

    function getCalendarItems() {
        const taskItems = state.tasks
            .filter(task => task.deadline)
            .map(task => ({
                ...task,
                isTask: true,
                start_time: task.deadline,
                end_time: null,
                title: task.title || "Untitled task",
                description: task.description || "",
            }));

        return [
            ...state.entries.map(entry => ({
                ...entry,
                isTask: false,
            })),
            ...taskItems,
        ];
    }

    function entriesOnDay(date) {
        const key = dateKey(date);

        return getCalendarItems()
            .filter(item => {
                const start = getEntryStart(item);
                return start && dateKey(start) === key;
            })
            .sort((a, b) => getEntryStart(a) - getEntryStart(b));
    }

    function sameDay(a, b) {
        return dateKey(a) === dateKey(b);
    }

    function makeElement(tag, className, text) {
        const element = document.createElement(tag);

        if (className) element.className = className;
        if (text !== undefined) element.textContent = text;

        return element;
    }

    function render() {
        view.replaceChildren();

        if (state.mode === "paper") {
            renderPaperCalendar();
        } else {
            renderTimetable();
        }
    }

    function renderTimetable() {
        const weekStart = startOfWeek(state.date);
        const weekEnd = new Date(weekStart);
        weekEnd.setDate(weekEnd.getDate() + 6);

        heading.textContent =
            `${formatDate(weekStart, { day: "numeric", month: "short" })} – ` +
            formatDate(weekEnd, { day: "numeric", month: "short", year: "numeric" });

        const wrapper = makeElement("div", "timetable-scroll");
        const grid = makeElement("div", "timetable-grid");

        // The first column contains the hours.
        grid.append(makeElement("div", "time-corner", ""));

        const days = [];

        for (let i = 0; i < 7; i++) {
            const day = new Date(weekStart);
            day.setDate(weekStart.getDate() + i);
            days.push(day);

            const header = makeElement("div", "timetable-day-header");
            const name = makeElement(
                "span",
                "timetable-day-name",
                formatDate(day, { weekday: "short" })
            );
            const number = makeElement(
                "span",
                "timetable-day-number",
                String(day.getDate())
            );

            if (sameDay(day, new Date())) {
                number.classList.add("is-today");
            }

            header.append(name, number);
            grid.append(header);
        }

        // Display 07:00–22:00.
        const firstHour = 7;
        const lastHour = 22;

        for (let hour = firstHour; hour < lastHour; hour++) {
            grid.append(
                makeElement(
                    "div",
                    "time-label",
                    `${String(hour).padStart(2, "0")}:00`
                )
            );

            for (const day of days) {
                const cell = makeElement("div", "timetable-cell");
                cell.dataset.date = dateKey(day);
                cell.dataset.hour = String(hour);

                grid.append(cell);
            }
        }

        wrapper.append(grid);
        view.append(wrapper);

        // Put each event into its weekday column and time position.
        for (const entry of getCalendarItems()) {
            const start = getEntryStart(entry);
            if (!start) continue;

            const dayIndex = days.findIndex(day => sameDay(day, start));
            if (dayIndex < 0) continue;

            const hour = start.getHours() + start.getMinutes() / 60;
            if (hour < firstHour || hour >= lastHour) continue;

            const end = getEntryEnd(entry);
            const duration = end
                ? Math.max(0.5, (end - start) / 3600000)
                : 1;

            const event = makeElement("button", "timetable-event");
            event.type = "button";

            if (entry.isTask) {
                event.classList.add("timetable-task");
            }

            event.title = entry.description || entry.title || "Event";

            event.append(
                makeElement(
                    "span",
                    "timetable-event-time",
                    formatDate(start, { hour: "2-digit", minute: "2-digit" })
                ),
                makeElement(
                    "span",
                    "timetable-event-title",
                    entry.title || "Untitled event"
                )
            );

            event.addEventListener("click", () => showEventDetails(entry));

            const rowHeight = 60;
            const top = (start.getMinutes() / 60) * rowHeight;
            const height = Math.max(30, duration * rowHeight);

            event.style.position = "absolute";
            event.style.top = `${top + 2}px`;
            event.style.height = `${height - 4}px`;
            event.style.left = "3px";
            event.style.right = "3px";

            const column = grid.querySelector(
                `.timetable-cell[data-date="${dateKey(days[dayIndex])}"][data-hour="${Math.floor(hour)}"]`
            );

            if (column) {
                column.append(event);
                column.classList.add("has-event");
            }
        }
    }

    function renderPaperCalendar() {
        const year = state.date.getFullYear();
        const month = state.date.getMonth();

        heading.textContent = formatDate(state.date, {
            month: "long",
            year: "numeric",
        });

        const grid = makeElement("div", "paper-calendar");

        for (const dayName of ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]) {
            grid.append(makeElement("div", "paper-weekday", dayName));
        }

        const firstDay = new Date(year, month, 1);
        const offset = (firstDay.getDay() + 6) % 7;
        const daysInMonth = new Date(year, month + 1, 0).getDate();

        // Complete the first week with the previous month's dates.
        for (let i = 0; i < offset; i++) {
            const date = new Date(year, month, 1 - offset + i);
            grid.append(makePaperDay(date, true));
        }

        for (let day = 1; day <= daysInMonth; day++) {
            grid.append(makePaperDay(new Date(year, month, day), false));
        }

        const totalCells = offset + daysInMonth;
        const trailing = (7 - (totalCells % 7)) % 7;

        for (let i = 1; i <= trailing; i++) {
            grid.append(makePaperDay(new Date(year, month + 1, i), true));
        }

        view.append(grid);
    }

    function makePaperDay(date, outsideMonth) {
        const cell = makeElement("button", "paper-day");
        cell.type = "button";

        if (outsideMonth) cell.classList.add("outside-month");
        if (sameDay(date, new Date())) cell.classList.add("is-today");
        if (state.selectedDay && sameDay(date, state.selectedDay)) {
            cell.classList.add("is-selected");
        }

        const number = makeElement("span", "paper-day-number", String(date.getDate()));
        cell.append(number);

        const events = entriesOnDay(date);

        if (events.length > 0) {
            const list = makeElement("span", "paper-day-events");

            for (const entry of events.slice(0, 2)) {
                list.append(
                    makeElement("span", "paper-event-dot", entry.title || "Event")
                );
            }

            if (events.length > 2) {
                list.append(
                    makeElement("span", "paper-more", `+${events.length - 2} more`)
                );
            }

            cell.append(list);
        }

        cell.addEventListener("click", () => {
            state.selectedDay = date;
            render();
            renderSelectedDay(date);
        });

        return cell;
    }

    function renderSelectedDay(date) {
        const card = document.querySelector("#selected-day-card");
        const title = document.querySelector("#selected-day-heading");
        const container = document.querySelector("#selected-day-events");

        card.hidden = false;
        title.textContent = formatDate(date, {
            weekday: "long",
            day: "numeric",
            month: "long",
            year: "numeric",
        });

        renderEventList(entriesOnDay(date), container);
    }

    function renderEventList(entries, container) {
        container.replaceChildren();

        if (entries.length === 0) {
            container.append(makeElement("p", "muted", "No events on this day."));
            return;
        }

        for (const entry of entries) {
            const row = makeElement("div", "task-row");
            const details = makeElement("div", "task-details");
            const start = getEntryStart(entry);
            const end = getEntryEnd(entry);

            details.append(
                makeElement("div", "task-title", entry.title || "Untitled event"),
                makeElement(
                    "div",
                    "task-meta",
                    start
                        ? `${formatDate(start, { hour: "2-digit", minute: "2-digit" })}` +
                          (end ? ` – ${formatDate(end, { hour: "2-digit", minute: "2-digit" })}` : "")
                        : "Time unavailable"
                )
            );

            if (entry.description) {
                details.append(
                    makeElement("p", "task-description", entry.description)
                );
            }

            row.append(details);

            if (entry.id) {
                const remove = makeElement("button", "small-button delete-button", "Delete");
                remove.type = "button";
                remove.addEventListener("click", () => deleteEntry(entry));
                row.append(remove);
            }

            container.append(row);
        }
    }

    function showEventDetails(entry) {
        const start = getEntryStart(entry);
        const end = getEntryEnd(entry);

        const dateText = start
            ? formatDate(start, {
                weekday: "long",
                day: "numeric",
                month: "long",
                hour: "2-digit",
                minute: "2-digit",
            })
            : "Time unavailable";

        const endText = end
            ? `\nEnds: ${formatDate(end, { hour: "2-digit", minute: "2-digit" })}`
            : "";

        window.alert(
            `${entry.title || "Untitled event"}\n\n${dateText}${endText}` +
            `${entry.description ? `\n\n${entry.description}` : ""}` +
            `${entry.task_id ? `\n\nLinked task: ${entry.task_id}` : ""}`
        );
    }

    async function loadEntries() {
        setMessage("Loading calendar...");

        try {
            const [plannerData, taskData] = await Promise.all([
                apiGetPlannerEntries(),
                apiGetTasks(),
            ]);

            state.entries = plannerData?.entries ?? [];
            state.tasks = taskData?.tasks ?? [];

            const taskCount = state.tasks.filter(task => task.deadline).length;
            const total = state.entries.length + taskCount;

            setMessage(
                `${state.entries.length} planner event(s) and ` +
                `${taskCount} task(s) with deadlines loaded.`
            );

            render();

            if (state.mode === "paper" && state.selectedDay) {
                renderSelectedDay(state.selectedDay);
            }
        } catch (error) {
            setMessage(
                `Couldn't load calendar data: ${error.message}`,
                "error"
            );

            console.error("Load calendar data:", error);
        }
    }

    async function deleteEntry(entry) {
        if (!window.confirm(`Delete "${entry.title || "this event"}"?`)) {
            return;
        }

        try {
            await apiDeletePlannerEntry(entry.id);
            state.entries = state.entries.filter(item => item.id !== entry.id);

            setMessage("Event deleted.", "success");
            render();

            if (state.selectedDay) {
                renderSelectedDay(state.selectedDay);
            }
        } catch (error) {
            setMessage(`Couldn't delete event: ${error.message}`, "error");
            console.error("Delete planner entry:", error);
        }
    }

    document.querySelector("#calendar-prev").addEventListener("click", () => {
        if (state.mode === "paper") {
            state.date.setMonth(state.date.getMonth() - 1);
        } else {
            state.date.setDate(state.date.getDate() - 7);
        }

        render();
    });

    document.querySelector("#calendar-next").addEventListener("click", () => {
        if (state.mode === "paper") {
            state.date.setMonth(state.date.getMonth() + 1);
        } else {
            state.date.setDate(state.date.getDate() + 7);
        }

        render();
    });

    document.querySelector("#calendar-today").addEventListener("click", () => {
        state.date = new Date();
        render();
    });

    document.querySelectorAll("[data-mode]").forEach(button => {
        button.addEventListener("click", () => {
            state.mode = button.dataset.mode;

            document.querySelectorAll("[data-mode]").forEach(item => {
                item.classList.toggle("active", item === button);
            });

            render();

            const selectedCard = document.querySelector("#selected-day-card");
            if (state.mode !== "paper") {
                selectedCard.hidden = true;
            } else if (state.selectedDay) {
                renderSelectedDay(state.selectedDay);
            }
        });
    });

    document.querySelector("#toggle-create-form").addEventListener("click", () => {
        formCard.hidden = !formCard.hidden;
    });

    document.querySelector("#cancel-create").addEventListener("click", () => {
        formCard.hidden = true;
    });

    form.addEventListener("submit", async event => {
        event.preventDefault();

        const button = form.querySelector('button[type="submit"]');
        const message = document.querySelector("#event-form-message");
        const formData = new FormData(form);

        const title = String(formData.get("title") ?? "").trim();
        const rawStart = String(formData.get("start_time") ?? "");
        const rawEnd = String(formData.get("end_time") ?? "");

        if (!title || !rawStart) {
            message.textContent = "Title and start time are required.";
            message.className = "form-message error";
            return;
        }

        const start = new Date(rawStart);
        const end = rawEnd ? new Date(rawEnd) : null;

        if (Number.isNaN(start.getTime()) || (end && Number.isNaN(end.getTime()))) {
            message.textContent = "Please enter valid dates.";
            message.className = "form-message error";
            return;
        }

        if (end && end <= start) {
            message.textContent = "The end time must be after the start time.";
            message.className = "form-message error";
            return;
        }

        const entry = {
            title,
            description: String(formData.get("description") ?? "").trim(),
            task_id: String(formData.get("task_id") ?? "").trim() || null,
            start_time: start.toISOString(),
            end_time: end ? end.toISOString() : null,
        };

        button.disabled = true;
        message.textContent = "Saving event...";
        message.className = "form-message";

        try {
            await apiCreatePlannerEntry(entry);

            form.reset();
            formCard.hidden = true;

            state.date = new Date(start);
            state.selectedDay = new Date(start);

            message.textContent = "Event created.";
            message.className = "form-message success";

            await loadEntries();

            if (state.mode === "paper") {
                renderSelectedDay(state.selectedDay);
            }
        } catch (error) {
            message.textContent = `Couldn't create event: ${error.message}`;
            message.className = "form-message error";
            console.error("Create planner entry:", error);
        } finally {
            button.disabled = false;
        }
    });

    loadEntries();
});