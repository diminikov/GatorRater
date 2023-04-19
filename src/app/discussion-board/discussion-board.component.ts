import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { DataService } from '../data.service';

@Component({
  selector: 'app-discussion-board',
  templateUrl: './discussion-board.component.html',
  styleUrls: ['./discussion-board.component.css']
})
export class DiscussionBoardComponent implements OnInit {
  className: string = '';
  classId: string = '';
  posts: string[] = [];
  newPost: string = '';

  constructor(private route: ActivatedRoute, private dataService: DataService) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      // Get classId parameter from URL
      this.classId = params.get('classId')!;
      // Get className from dataService
      this.className = this.dataService.getClass(this.classId).name;
      // Get posts from dataService
      this.posts = this.dataService.getClass(this.classId).posts;
    });
  }

  onSubmit(classId: string, post: string) {
    this.dataService.addPost(classId, post).subscribe((response) => {
      console.log('Post added successfully:', response);
    }, (error) => {
      console.error('Error adding post:', error);
    });
  }

}

