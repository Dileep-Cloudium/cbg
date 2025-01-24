import { ErrorHandler, Injectable } from '@angular/core';
import { ResponseModel } from '../../shared/response.model';

@Injectable()
export class GlobalErrorHandler implements ErrorHandler {

    handleError(error: ResponseModel['data']): void {
        const chunkFailedMessage = /Loading chunk [\d]+ failed/;
        if (chunkFailedMessage.test(error.message)) {
            window.location.reload();
        }
    }
}