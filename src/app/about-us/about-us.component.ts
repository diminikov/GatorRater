import { Component } from '@angular/core';
import { AboutUs } from '../aboutUs.model';

@Component({
  selector: 'app-about-us',
  templateUrl: './about-us.component.html',
  styleUrls: ['./about-us.component.css']
})
export class AboutUsComponent {
  persons : AboutUs [] = [
    new AboutUs('Michael Sawarynski', 'Hello, I am from Gulf breeze Fl.', 'Front-End Development'), 
    new AboutUs('Yonas', 'Hello, I am from Jax.', 'Front-End Development'), 
    new AboutUs('Nish', 'Hello, I am from Australia.', 'Back-End Development'), 
    new AboutUs('Dimitri', 'Hello, I am from Gulf breeze Fl.', 'Front-End Development') 

  ];

}
