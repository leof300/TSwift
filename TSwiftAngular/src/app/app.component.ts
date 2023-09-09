import { Component, OnInit, OnDestroy } from '@angular/core';
import { CodeModel } from '@ngstack/code-editor';
import { GoApiService } from './service/go-api.service';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent{

  theme = 'vs-dark';
  editorValue = "";
  consoleValues: any = [];
  simbolValues: any = [];
  errorValues: any = [];





  constructor(private apiGo: GoApiService){}

  codeModel: CodeModel = {
    language: 'golang',
    uri: 'main.json',
    value: ''
  };

  options = {
    contextmenu: true,
    minimap: {
      enabled: true
    }
  };

  onCodeChanged(value: any) {
    this.editorValue = value;
  }


  sendText() {
    this.consoleValues = [];
    this.simbolValues = [];
    this.errorValues = [];

    // const body = {"InputText":JSON.stringify(this.editorValue)}
    this.apiGo.execute({"InputText":this.editorValue})
      .subscribe((data: any) => {
        console.log(data)
        if(data){
          if(data.Console){
            this.consoleValues = data.Console;
          }
          if(data.Exceptions){
            this.errorValues = data.Exceptions;
          }
          if(data.TDS){
            this.simbolValues = data.TDS;
          }
        }
    });
  }

  ast(){
    this.apiGo.tree()
      .subscribe(data => {
        console.log(data);
      })
  }

}
