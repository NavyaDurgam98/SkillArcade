import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css'],
  standalone: true,
  imports: [CommonModule]
})
export class DashboardComponent {
  technologies = [
    { name: 'Angular', description: 'A powerful frontend framework by Google.', image:'/assets/angular.png' },
    { name: 'React', description: 'A library for building user interfaces by Facebook.', image: 'assets/react.png' },
    { name: 'Vue.js', description: 'A progressive JavaScript framework.', image: 'assets/vue.png' },
    { name: 'Node.js', description: 'A runtime for executing JavaScript on the server.', image: 'assets/node.png' },
    { name: 'Python', description: 'A versatile programming language used in many fields.', image: 'assets/python.png' },
    { name: 'Java', description: 'A widely-used programming language for applications.', image: 'assets/java.png' },
    { name: 'C++', description: 'A powerful language for system and application development.', image: 'assets/cpp.png' },
    { name: 'SQL', description: 'A language for managing relational databases.', image: 'assets/sql.png' }
  ];
}
