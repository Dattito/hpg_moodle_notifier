import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor,
  HttpErrorResponse
} from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { AlertService } from './alert.service';

@Injectable()
export class HttpErrorInterceptor implements HttpInterceptor {

  constructor(
    private alert: AlertService
  ) {}

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    return next.handle(request)
     .pipe(
       catchError((err: HttpErrorResponse) => {

        let errorMessage = '';
        if (err.error instanceof ErrorEvent) {
          // client-side error
          errorMessage = `Etwas ist schiefgelaufen.`;
        } else {
          // server-side error
          errorMessage = err.error["msg"];
        }
        this.alert.error(errorMessage);

        const error = err.error?.message || err.statusText;
        return throwError(error);
       })
     );
  }
}
