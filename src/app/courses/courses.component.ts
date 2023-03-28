import { Component,Output, EventEmitter } from '@angular/core';
import { Course } from './course.model';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';


@Component({
  selector: 'app-courses',
  templateUrl: './courses.component.html',
  styleUrls: ['./courses.component.css']
})
export class CoursesComponent {
  constructor(private httpClient: HttpClient, router:Router){

  }

  courses : Course [] = [
    new Course('CEN3031', 'Introduction to Software Engineering'),
    new Course('COP3503C', 'Programming Fundamentals 2'),
    new Course('COP3502C', 'Programming Fundamentals 1'),
    new Course('CDA3101', 'Introduction to Computer Organization'), 
    new Course('COP3530', 'Data Structures and Algorithm'),
    new Course('COT3100', 'Applications of Discrete Structures'),
    new Course('CIS4301', 'Information and Database Systems 1'),
    new Course('CIS4914', 'Senior Project'),
    new Course('CNT4007', 'Computer Network Fundamentals'),
    new Course('COP4020', 'Programming Language Concepts'),
    new Course('COP4533', 'Algorithm Abstraction and Design'),
    new Course('COP4600', 'Operating Systems'),
    new Course('EGS4034', 'Engineering Ethics and Professionalism'),

  ];




  courseClicked(course: Course){
    console.log(Course.name);

  }

  onSelected(index: number){
    //this.courseSelected.emit();
    //console.log
    return index;
  }

  searchTerm = '';
  classes: Course[] = [];
  term = '';
  allClasses: Course[] = [];


  
  ngOnInit(): void {
    this.httpClient.get<Course[]>('').subscribe((data: Course[]) => {
      this.classes = data;
      this.allClasses = this.classes;
    });

  }

  search(value:string): void{
    this.classes = this.allClasses.filter((val)=>
    val.name.toLowerCase().includes(value)
    );
  }


}
