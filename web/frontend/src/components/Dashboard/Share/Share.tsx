import React from 'react';
import { FacebookIcon, TwitterIcon, LinkedinIcon, PinterestIcon } from 'react-share';

const Share: React.FC = () => {
    function shareOnFB() {
        window.open('https://www.facebook.com/sharer/sharer.php?u=progflow.tech')
    }

    function shareOnTwitter() {
        window.open('https://twitter.com/intent/tweet?text=Try%20out%20this%20awesome%20tool!%0A%0Ahttps%3A//progflow.tech')
    }

    function shareOnLinkedIn() {
        window.open('https://www.linkedin.com/shareArticle?mini=true&url=https%3A//progflow.tech')
    }

    function shareOnPintrest() {
        window.open('https://pinterest.com/pin/create/button/?url=https%3A//progflow.tech/&media=')
    }

    return (
        <div className='flex justify-between cursor-pointer align-center h-[7.6rem] items-center'>
            <FacebookIcon size={40} borderRadius={5} onClick={shareOnFB} />
            <TwitterIcon size={40} borderRadius={5} onClick={shareOnTwitter} />
            <LinkedinIcon size={40} borderRadius={5} onClick={shareOnLinkedIn} />
            <PinterestIcon size={40} borderRadius={5} onClick={shareOnPintrest} />
        </div>
    );
};

export default Share
