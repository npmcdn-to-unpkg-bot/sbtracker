const host = 'localhost:8080'

var SandboxList = React.createClass({
    getInitialState: function() {
        return {data: []};
    },

    componentDidMount: function() {
        var s = this;
        fetch('http://localhost:8080/sandbox/list.json')
        .then(function (response) {
            if(response.ok) {
                response.json().then(function(json) {
                    s.setState({data: json});
                });
            } else {
                console.log('Network response was not ok.');
            }
        });
    },

    render: function() {
        console.log(this.state);
        var sandboxes = this.state.data.map(function(sandbox) {
            return (
                <Sandbox
                 key={sandbox.Id}
                 id={sandbox.Id}
                 url={sandbox.Url}
                 owner={sandbox.Owner}
                 branch={sandbox.Branch}/>
            );
        });
        return (
            <div className="sandbox-list">
                {sandboxes}
            </div>
        );
    }
});

var Sandbox = React.createClass({
    propTypes: {
        id: React.PropTypes.number.isRequired,
        url: React.PropTypes.string.isRequired,
        owner: React.PropTypes.string,
        branch: React.PropTypes.string
    },

    render: function() {
        return (
            <div key="{this.prop.id}" className="sandbox">
                <span className="prop-id">{this.props.id}</span>
                <span className="prop-url">
                    <a href={this.props.url} target="_blank">
                        <i className="fa fa-external-link"></i>
                    </a>
                </span>
                <div>
                    <span className={this.props.owner ? 'label active' : 'label'}>sisu</span>
                    <select className="prop-owner">
                        <option value={this.props.owner} selected>{this.props.owner}</option>
                    </select>
                    <span className="prop-branch">{this.props.branch}</span>
                </div>
                <div>
                    <span className={this.props.owner ? 'label active' : 'label'}>wl</span>
                    <select className="prop-owner">
                        <option value={this.props.owner} selected>{this.props.owner}</option>
                    </select>
                    <span className="prop-branch">{this.props.branch}</span>
                </div>
                <div>
                    <span className={this.props.owner ? 'label active' : 'label'}>ms</span>
                    <select className="prop-owner">
                        <option value={this.props.owner} selected>{this.props.owner}</option>
                    </select>
                    <span className="prop-branch">{this.props.branch}</span>
                </div>
            </div>
        );
    }
})

ReactDOM.render(
    <div>
        <h1>Sandboxes</h1>
        <SandboxList url="http://localhost:8080/sandbox/list.json"></SandboxList>
    </div>,
    document.getElementById('app')
);