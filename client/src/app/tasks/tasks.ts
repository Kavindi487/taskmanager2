import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Task, User, TaskService } from '../task.service';

@Component({
  selector: 'app-tasks',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './tasks.html',
  styleUrl: './tasks.css'
})
export class TasksComponent implements OnInit {
  tasks: Task[] = [];
  users: User[] = [];
  newTitle = '';
  selectedUserId: number | null = null;  // no user selected by default

  constructor(private taskService: TaskService) {}

  ngOnInit() {
    this.taskService.getUsers().subscribe(users => {
      this.users = users ?? [];
    });
  }

  onUserSelect(userId: number) {
    this.selectedUserId = userId;
    this.loadTasks();
  }

  loadTasks() {
    if (!this.selectedUserId) return;
    this.taskService
      .getTasks(this.selectedUserId)
      .subscribe(tasks => this.tasks = tasks ?? []);
  }

  addTask() {
    if (!this.newTitle.trim() || !this.selectedUserId) return;
    this.taskService
      .createTask(this.newTitle, this.selectedUserId)
      .subscribe(() => {
        this.newTitle = '';
        this.loadTasks();
      });
  }

  complete(task: Task) {
    this.taskService
      .markDone(task.ID)
      .subscribe(() => task.Done = true);
  }

  getSelectedUserName(): string {
  const user = this.users.find(u => u.ID === this.selectedUserId);
  return user ? user.Name : '';
  }
}