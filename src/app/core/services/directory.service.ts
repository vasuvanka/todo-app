import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { Directory } from '../models/directory';
import { GeneralResponse } from '../models/user';

@Injectable({
  providedIn: 'root'
})
export class DirectoryService {

  constructor(private http: HttpClient) { }

  getDirectories(id: string): Promise<Directory[]> {
    return this.http.get<Directory[]>(`${environment.api}/directories/${id}/subdirs`).toPromise()
  }

  createDirectory(dir: Directory): Promise<Directory> {
    return this.http.post<Directory>(`${environment.api}/directories`, dir).toPromise()
  }

  updateDirectory(dir: Directory): Promise<GeneralResponse> {
    return this.http.put<GeneralResponse>(`${environment.api}/directories/${dir.id}`, dir).toPromise()
  }

  deleteDirectory(id: string): Promise<GeneralResponse> {
    return this.http.delete<GeneralResponse>(`${environment.api}/directories/${id}`).toPromise()
  }

}
