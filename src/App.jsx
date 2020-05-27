import React from 'react';
import './App.css';
import Transactions from "./components/Transaction";
import Balance from "./components/Balance";
import { getTransactions, getBalance } from "./request/request";

class App extends React.PureComponent {
  constructor(props) {
    super(props);
    this.state = {
      balance: 0,
      transactions: [],
      errorMessage: ''
    };
    this.getTransactions = this.getTransactions.bind(this);
    this.getBalance = this.getBalance.bind(this);
    this.getData = this.getData.bind(this);
    this.refresh = this.refresh.bind(this);
  }

  async componentDidMount() {
    await this.getData();
  }

  getData = async() => {
    try {
      await this.getTransactions();
      await this.getBalance();
    } catch (error) {
      console.log({error})
      this.setState({
        errorMessage: error.message
      });
    }
  }

  getTransactions = async() => {
    const transactions = await getTransactions()
    this.setState({
      transactions: transactions || []
    });
  }

  getBalance = async() => {
    const balance = await getBalance()
    this.setState({
        balance: balance
    });
  }

  refresh = async () => await this.getData();

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <div className="wrapper">
            <h1>Transactions</h1>
            <button className="refresh" onClick={this.refresh}>Refresh</button>
          </div>
          <label className="error">{this.state.errorMessage}</label>
          <Transactions transactions={this.state.transactions} />
          <Balance balance={this.state.balance} />
        </header>
      </div>
      );
  }
}

export default App;