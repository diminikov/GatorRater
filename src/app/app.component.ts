import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  public user = '';
  public password = '';

  constructor(
    private httpClient: HttpClient
      ){}


loadedFeature = 'course';

onNavigate(feature: string){
  this.loadedFeature = feature;
}

  title = 'GatorRater';




}
