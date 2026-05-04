import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

// Model — matches your Go Task struct
export interface Task {
  id: number;
  title: string;
  done: boolean;
}

@Injectable({ providedIn: 'root' })
export class TaskService {
  private base = 'http://localhost:8080';

  // inject HttpClient via constructor
  constructor(private http: HttpClient) {}

  // GET /tasks?user=1
  getTasks(userId: number): Observable<Task[]> {
    return this.http
      .get(`${this.base}/tasks`, {
        params: new HttpParams().set('user', userId),
        responseType: 'text'
      })
      .pipe(
        map(raw => raw.trim().split('\n')
          .filter(t => t)
          .map((title, i) => ({ id: i + 1, title, done: false }))
        )
      );
  }

  // POST /tasks  (form-data, NOT JSON!)
  createTask(title: string): Observable<string> {
    const body = new HttpParams().set('title', title);
    return this.http.post(`${this.base}/tasks`, body.toString(), {
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
      responseType: 'text'
    });
  }

  // GET /tasks/done?id=2
  markDone(id: number): Observable<string> {
    return this.http.get(`${this.base}/tasks/done`, {
      params: new HttpParams().set('id', id),
      responseType: 'text'
    });
  }
}