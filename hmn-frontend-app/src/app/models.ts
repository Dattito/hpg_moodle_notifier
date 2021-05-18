export interface PostSignalVerificationResponse {
    id: string
}

export interface PostAssignmentResponse {
    msg: string
}

export interface GetMoodleTokenResponse {
    token: string
}

export class Alert {
    id?: string;
    type?: AlertType;
    message?: string;
    autoClose?: boolean;
    keepAfterRouteChange?: boolean;
    fade?: boolean;

    constructor(init?:Partial<Alert>) {
        Object.assign(this, init);
    }
}

export enum AlertType {
    Success,
    Error,
    Info,
    Warning
}