import { Component, EventEmitter, Output,Input } from '@angular/core';
import { Course } from '../course.model';
import { CoursesComponent } from '../courses.component';

@Component({
  selector: 'app-course-details',
  templateUrl: './course-details.component.html',
  styleUrls: ['./course-details.component.css']
})
export class CourseDetailsComponent{
  @Output() courseSelected = new EventEmitter<Course>();



  onSelected(index: number){
    //this.courseSelected.emit();
    console.log(index);
  }

}
