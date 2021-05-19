import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable} from 'rxjs';
import { GetMoodleTokenResponse, PostAssignmentResponse, PostSignalVerificationResponse } from './models';

@Injectable({
  providedIn: 'root'
})
export class HmnBackendService {

  readonly hmnBackendUrl = "https://hmn.dattid.to/api"

  constructor(private http: HttpClient) { }

  getMoodleToken(username: string, password: string): Observable<GetMoodleTokenResponse> {
    return this.http.post<GetMoodleTokenResponse>(this.hmnBackendUrl + '/v1/moodleToken', {
      username,
      password
    });
  }

  getSignalVerificationCode(moodleToken: string, phoneNumber: string): Observable<PostSignalVerificationResponse> {
    return this.http.post<PostSignalVerificationResponse>(this.hmnBackendUrl + '/v1/signalVerifications', {
      moodleToken,
      phoneNumber
    });
  }

  registerAssignments(verificationID: string, verificationCode: string): Observable<PostAssignmentResponse> {
    return this.http.post<PostAssignmentResponse>(this.hmnBackendUrl + '/v1/assignments', {
      id: verificationID,
      verificationCode
    });
  }

  deleteAssignments(verificationID: string, verificationCode: string): Observable<PostAssignmentResponse> {
    //let httpParams = new HttpParams().set('id', verificationID)
    //httpParams.set('verificationCode', verificationCode)
    return this.http.request<PostAssignmentResponse>('delete', this.hmnBackendUrl + '/v1/assignments', {
      body: {
        id: verificationID,
        verificationCode
      }
    });
  }
}
