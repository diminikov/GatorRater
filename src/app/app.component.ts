import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';


interface someClass {
  name: string;
  professor: string;
}
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
  title = 'GatorRater';

  searchTerm = '';
  classes: someClass[] = [];
  term = '';
  allClasses: someClass[] = [];

  constructor(private http: HttpClient) { }
  
  ngOnInit(): void {
    this.http.get<someClass[]>('./assets/data/classes.json')
    .subscribe((data: someClass[]) => {
      this.classes = data;
      this.allClasses = this.classes;
    });
  }

  search(value: string): void{
    this.classes = this.allClasses.filter((val) =>
    val.name.toLowerCase().includes(value)
    );
  }

}
