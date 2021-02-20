import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Todo } from 'src/app/core/models/todo';
import { GeneralResponse } from 'src/app/core/models/user';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class TodoService {

  constructor(private http: HttpClient) { }

  getTodos(dirId: string): Promise<Todo[]> {
    return this.http.get<Todo[]>(`${environment.api}/directories/${dirId}/todos`).toPromise()
  }

  getTodoById(id: string): Promise<Todo> {
    return this.http.get<Todo>(`${environment.api}/todos/${id}`).toPromise()
  }

  createTodo(todo: Todo): Promise<Todo> {
    return this.http.post<Todo>(`${environment.api}/todos`, todo).toPromise()
  }

  updateTodo(todo: Todo): Promise<GeneralResponse> {
    return this.http.put<GeneralResponse>(`${environment.api}/todos/${todo.id}`, todo).toPromise()
  }

  removeTodo(id: string): Promise<GeneralResponse> {
    return this.http.delete<GeneralResponse>(`${environment.api}/todos/${id}`).toPromise()
  }
}
