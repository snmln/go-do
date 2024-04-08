import { Component, OnInit } from '@angular/core';
import { TodoService, Todo } from '../todo.service';

@Component({
  selector: 'app-todo',
  templateUrl: './todo.component.html',
  styleUrls: ['./todo.component.scss']
})
export class TodoComponent implements OnInit {

  activeTodos: Todo[];
  completedTodos: Todo[];
  todoMessage: string;


  constructor(private todoService: TodoService) { }

  ngOnInit() {
    this.getAll();
  }

  getAll() {
    this.todoService.getTodoList().subscribe(
      (data) => {
        if (Array.isArray(data)) {
          this.activeTodos = data.filter((a: any) => !a.complete);
          this.completedTodos = data.filter((a: any) => a.complete);
          return;
        }
      }
    );
  }

  addTodo() {
    var newTodo: Todo = {
      message: this.todoMessage,
      id: '',
      complete: false
    };

    this.todoService.addTodo(newTodo).subscribe(() => {
      this.getAll();
      this.todoMessage = '';
    });
  }

  completeTodo(todo: Todo) {
    this.todoService.completeTodo(todo).subscribe(() => {
      this.getAll();
    });
  }

  revertTodo(todo: Todo) {
    // console.log("this is from revert", todo)
    // var updatedTodo = { id: todo.id, message: todo.message, complete: false }
    this.todoService.revertTodo(todo).subscribe(() => {
      this.getAll();
    });
  }

  deleteTodo(todo: Todo) {
    this.todoService.deleteTodo(todo).subscribe(() => {
      this.getAll();
    })
  }
}