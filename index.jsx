class Link extends React.Component {
    constructor(props) {
        super(props)
        console.log(props); 
    }

    render () {
        return (
            <a href={this.props.to}>{this.props.children}</a>
        );
    }
}


// import Menu from 'antd/lib/menu';  // 加载 JS
// import 'antd/antd/lib/menu/style/css';        // 加载 CSS
// import Icon from 'antd/lib/icon';  // 加载 JS
// import 'antd/lib/icon/style/css';  // 加载 JS

// 左侧菜单栏
const SubMenu = antd.Menu.SubMenu;
const Menu = antd.Menu;
const Icon = antd.Icon;

class HomeLayout extends React.Component {
    render () {
        const {children} = this.props;
        return (
            <div>
                <header className="header">
                        <Link to="/">ReactManager</Link>
                </header>
        
                <main className="main">
                    <div className="menu">
                    <Menu mode="inline" theme="dark" style={{width: '240'}}>
                        <SubMenu key="user" title={<span><Icon type="user"/><span>用户管理</span></span>}>
                        <Menu.Item key="user-list">
                            <Link to="/user/list">用户列表</Link>
                        </Menu.Item>
                        <Menu.Item key="user-add">
                            <Link to="/user/add">添加用户</Link>
                        </Menu.Item>
                        </SubMenu>
            
                        <SubMenu key="book" title={<span><Icon type="book"/><span>图书管理</span></span>}>
                        <Menu.Item key="book-list">
                                <Link to="/book/list">图书列表</Link>
                        </Menu.Item>
                        <Menu.Item key="book-add">
                                <Link to="/book/add">添加图书</Link>
                        </Menu.Item>
                        </SubMenu>
                    </Menu>
                    </div>
            
                    <div className="content" id="content">
                            <Home />
                    </div>
                </main>
            </div>
        );
    }
}


class Home extends React.Component {
// 构造器
constructor(props) {
    super(props);
    // 定义初始化状态
    this.state = {};
}

render() {
    return (
    <div className="welcome">
        Welcome
    </div>
    );
}
}    

ReactDOM.render(<HomeLayout />, mountNode);