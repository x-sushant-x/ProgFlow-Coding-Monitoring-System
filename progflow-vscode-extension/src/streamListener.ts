import { Readable } from 'stream'

export class DataStream extends Readable {
    private isEnded: boolean

    constructor() {
        super()
        this.isEnded = false
    }

    _read() {}

    public addData(data: string) {
        if (this.isEnded) {
            throw new Error('Stream has already ended. Cannot add more data.');
        }
        this.push(data);
    }
}