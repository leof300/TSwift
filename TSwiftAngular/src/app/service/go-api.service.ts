import {HttpClient, HttpHeaders} from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class GoApiService {

  constructor(private http: HttpClient) { }



  execute(text: any) {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
      })
    };

    return this.http.post('http://localhost:8080/tscompiler/ast',text, httpOptions);
  }

  tree(){
    return this.http.post("http://localhost:8080/tscompiler/tree", {});
  }

}
