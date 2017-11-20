import React, { Component } from 'react';
//import "node_modules/video-react/dist/video-react.css";
import videojs from 'video.js';
import Header from './Header';

class HomeComponent extends Component {
    render() {
        return (
            <div></div>
        );
    }
};

/*const videoJsOptions = {
    autoplay: true,
    controls: false,
    sources: [{
      src: 'http://localhost:3000/webcam/stream',
      type: 'video/mp4'
    }]
}

class HomeComponent extends Component {
    componentDidMount() {
        // instantiate video.js
        this.player = videojs(this.videoNode, videoJsOptions, function onPlayerReady() {
          console.log('onPlayerReady', this)
        });
    }

    componentWillUnmount() {
        if (this.player) {
            this.player.dispose()
        }
    }
    
    // wrap the player in a div with a `data-vjs-player` attribute
    // so videojs won't create additional wrapper in the DOM
    // see https://github.com/videojs/video.js/pull/3856
    render() {
        return (
        <div data-vjs-player>
            <video ref={ node => this.videoNode = node } className="video-js"></video>
        </div>
        )
    }
};*/

export default HomeComponent;