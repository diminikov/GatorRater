import { Component } from '@angular/core';
import { AboutUs } from '../aboutUs.model';

@Component({
  selector: 'app-about-us',
  templateUrl: './about-us.component.html',
  styleUrls: ['./about-us.component.css']
})
export class AboutUsComponent {
  persons : AboutUs [] = [
    new AboutUs('Michael Sawarynski', 'Hello, I am from Gulf breeze Fl.', 'Front-End Development','assets/img/MichaelSawarynskiHeadshot.jpeg'),
    new AboutUs('Yonas', 'Hello, I am from Jax.', 'Front-End Development','assets/img/YonasWorku.jpeg'),
    new AboutUs('Nish', 'Hello, I am from Australia.', 'Back-End Development','assets/img/nish.jpeg'),
    new AboutUs('Dimitri', 'Hello, I am from Gulf breeze Fl.', 'Front-End Development','assets/img/Dmitri.jpeg' )

  ];

}
