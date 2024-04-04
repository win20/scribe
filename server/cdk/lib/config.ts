import * as dotenv from 'dotenv';
import * as path from 'path';

dotenv.config({path: path.resolve(__dirname, '../.env')});

export type ConfigProps = {
    APP_NAME: string;
    REGION: string;
    ACCOUNT_NUMBER: string;
};

export const getConfig = (): ConfigProps => ({
    APP_NAME: 'scribe',
    REGION: process.env.REGION || 'eu-west-1',
    ACCOUNT_NUMBER: process.env.ACCOUNT_NUMBER || ''
});

