import superagent from "superagent";
import {URL_API} from "../constants/URLConstants";

export const getTransactions = async () => {
    try {
        const res = await superagent.get(`${URL_API}transactions`);
        return res.body;
    } catch(err) {
        throw err
    }
}

export const getBalance = async () => {
    try {
        const res = await superagent.get(`${URL_API}balance`);
        return res.body;
    } catch(err) {
        throw err
    }
}