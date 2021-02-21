import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-default-directory',
  templateUrl: './default-directory.component.html',
  styleUrls: ['./default-directory.component.css']
})
export class DefaultDirectoryComponent {

  constructor(private router: Router, private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.router.navigate(['.', '0'], { relativeTo: this.route })
  }

}
