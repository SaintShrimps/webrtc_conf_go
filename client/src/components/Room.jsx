import React, {useEffect, useRef} from 'react';
import { useLocation, useParams } from 'react-router';

//2.15 Session 2 : Building Video Chat Apps using WebRTC and Golang
const Room = () => {
    const { roomID } = useParams();
    //const params = useLocation(); //no
    const userVideo = useRef();
    const userStream = useRef();
    const partnerVideo = useRef();
    const peerRef = useRef();
    const webSocketRef = useRef();
   
    const openCamera = async () => {
        const allDevices = await navigator.mediaDevices.enumerateDevices();
        const cameras = allDevices.filter(
            (device) => device.kind == "videoinput"
        );
        console.log(cameras);

        const constraints = {
            audio: true,
            video: {
                deviceID: cameras[1].deviceId,
            },
        };

        try {
            return await navigator.mediaDevices.getUserMedia(constraints);
        } catch (error) {
            console.log(err);
        }
    }

    useEffect(()=>{
        openCamera().then((stream) => {
            userVideo.current.srcObject = stream;
            userStream.current = stream;

            webSocketRef.current = new WebSocket(
                `ws://localhost:8080/join?roomID=${roomID}`
            );

            webSocketRef.current.addEventListener("open", () => {
                webSocketRef.current.send(JSON.stringify({ join: true }));
            });

            webSocketRef.current.addEventListener("message", async (e) => {
                const message = JSON.parse(e.data)
                if (message.join) {
                    callUser();
                }

                if (message.offer) {
                    handleOffer(message.offer);
                }

                if (message.answer){
                    console.log("Receving Answer");
                    peerRef.current.setRemoteDescription(
                        new RTCSessionDescription(message.answer)
                    );
                }

                if (message.iceCandidate) {
                    console.log("Receiving and Adding ICE Candidate");
                    try {
                        await peerRef.current.addIceCandidate(message.iceCandidate);
                    } catch (error) {
                        console.log("Error Receiving ICE Candidate ", error);
                    }
                }

            });

        });

    });

    const callUser = () => {
        console.log("Calling Other User");
        peerRef.current = createPeer();

        userStream.current.getTracks().forEach((track) => {
            peerRef.current.addTrack(track, userStream.current);
        }); 
    };

    const handleOffer = async (offer) => {
        console.log("Received Offer Creating Answer");

        peerRef.current = createPeer();

        await peerRef.current.setRemoteDescription(
            new RTCSessionDescription(offer)
        );

        userStream.current.getTracks().forEach((track) => {
            peerRef.current.addTrack(track, userStream.current);
        }); 

        const answer = await peerRef.current.createAnswer();
        await peerRef.current.setLocalDescription(answer);

        webSocketRef.current.send(
            JSON.stringify(
                { answer: peerRef.current.localDescription }
            )
        );

    };

    const createPeer = () => {
        console.log("Creating Peer Connection");
        const peer = new RTCPeerConnection({
            iceServers: [{urls: "stun:stun.l.google.com:19302"},
            {url:'stun:stun1.l.google.com:19302'},
            {url:'stun:stun2.l.google.com:19302'},
            {url:'stun:stun3.l.google.com:19302'},
            {url:'stun:stun4.l.google.com:19302'},
        
        
        
        ],
        });

        peer.onnegotiationneeded = handleNegotiationNeeded;
        peer.onicecandidate = handleIceCandidateEven;
        peer.ontrack = handleTrackEvent;

        return peer;
    };

    const handleNegotiationNeeded = async () => {
        console.log("Creating Offer");

        try {
            const myOffer = await peerRef.current.createOffer();
            await peerRef.current.setLocalDescription(myOffer);

            webSocketRef.current.send(
                JSON.stringify({ offer: peerRef.current.localDescription })
            );

        } catch (error) {
            console.log("Error Creating Offer: ", error)
        }
    };

    const handleIceCandidateEven = (e) => {
        console.log("Found Ice Candidate");
        if (e.candidate) {
            console.log(e.candidate);
            webSocketRef.current.send(JSON.stringify({iceCandidate: e.candidate}));
        }
    };

    const handleTrackEvent = (e) => {
        console.log("Received Tracks");
        console.log("THIS FUCKING STREAM", e.streams);
        partnerVideo.current.srcObject = e.streams[0];
        
    };

    return (
        <div>
            <video autoPlay controls={true} ref={userVideo}> </video>
            <video autoPlay controls={true} ref={partnerVideo} ></video>
        </div>
        );
};

export default Room