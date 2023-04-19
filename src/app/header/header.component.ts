import { Component, EventEmitter,Output } from '@angular/core';
import { DataService } from '../data.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {

  constructor(private dataService: DataService, private router: Router){}
@Output() featureSelected = new EventEmitter<string>();

onSelect(feature: string){
  this.featureSelected.emit(feature);
}

onClick(){
  this.dataService.logout;
  this.router.navigate(['/signup']);
}

}
