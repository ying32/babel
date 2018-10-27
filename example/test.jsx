           //import { Button } from 'antd';

		   class App extends React.Component {
			constructor() {
				super()
				this.state = {
					loading: false,
					iconLoading: false,
				}

				this.enterLoading = () =>  {
					this.setState({ loading: true });
				}

				this.enterIconLoading = () =>  {     
					this.setState({ iconLoading: true });
				}
			}

			render() {
				return (
					<span>
						<Button type="primary" loading>
						Loading
						</Button>
						<Button type="primary" size="small" loading>
						Loading
						</Button>
						<br />
						<Button type="primary" loading={this.state.loading} onClick={this.enterLoading}>
						Click me!
						</Button>
						<Button type="primary" icon="poweroff" loading={this.state.iconLoading} onClick={this.enterIconLoading}>
						Click me!
						</Button>
						<br />
						<Button shape="circle" loading />
						<Button type="primary" shape="circle" loading />
					</span>
				);
			}
		}
	   ReactDOM.render(<App />, document.getElementById('app'));