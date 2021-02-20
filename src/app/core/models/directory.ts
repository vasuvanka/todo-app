export class Directory {
    title: string
    removed: boolean
    id: string
    when: Date
    parentId: string

    static fromJson(json: any = {}): Directory {
        const dir = new Directory()
        dir.title = json.title
        dir.removed = json.removed
        dir.id = json.id
        dir.parentId = json.parentId
        dir.when = json.when
        return dir
    }
}


