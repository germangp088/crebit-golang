import superagent from "superagent";
import {URL_API} from "../constants/URLConstants";

export const getTransactions = async () => {
    try {
        const res = await superagent.get(`${URL_API}transactions`);
        console.log(JSON.parse(res.text))
        return JSON.parse(res.text);
    } catch(err) {
        throw err
    }
}

export const getBalance = async () => {
    try {
        const res = await superagent.get(`${URL_API}balance`);
        console.log(JSON.parse(res.text))
        return JSON.parse(res.text);
    } catch(err) {
        throw err
    }
}