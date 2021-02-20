export type TodoType = "self" | "shared"
export type TodoPriority = "low" | "medium" | "high"
export type TodoStatus = "created" | "inprogress" | "completed"

export class Todo {
    title: string
    description: string
    type: TodoType
    priority: TodoPriority
    status: TodoStatus
    when: Date
    dueDate: Date
    id: string
    dirId: string

    static fromJson(json: any = {}): Todo {
        const todo = new Todo()
        todo.description = json.description
        todo.id = json.id
        todo.dueDate = new Date(json.dueDate)
        todo.priority = json.priority
        todo.status = json.status
        todo.title = json.title
        todo.type = json.type
        todo.when = new Date(json.when)
        todo.dirId = json.dirId || "0"
        return todo
    }
}
