import ChatHistory from "./components/Header/ChatHistory/ChatHistory";
import "./App.css";
import {connect, sendMsg} from "./api/";
import { Component } from "react";
import Header from "./components/Header/Header";

class App extends Component {
	constructor(props) {
		super(props);
		this.state = {
			chatHistory : []
		}

	}
	componentDidMount() {
		connect((msg) => {
			console.log("New Message")
			this.setState(prevState => ({
				chatHistory: [...this.state.chatHistory, msg]
			}))
			console.log(this.state);
		});
	}

	send() {
		console.log("Hello");
		sendMsg("Hello");
	}

	render() {
		return (
			<div className="App">
				<Header />
				<ChatHistory chatHistory={this.state.chatHistory} />
				<button onClick={this.send}>Hit</button>
			</div>
		);
	}
}


export default App;