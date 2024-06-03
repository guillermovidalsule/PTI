# Software

## Introduction

The software of our surveillance robot is designed to be user-friendly and efficient, allowing for seamless integration with various devices and systems. It is built using a combination of open-source software and hardware components, ensuring a robust and reliable platform for surveillance tasks.

In this section, we will delve into the details of our software, exploring its architecture, components, and functionalities. We will also discuss the key features and benefits of our software, providing a comprehensive overview of its capabilities and potential applications.

## Architecture

Our surveillance robot is built upon a modular and scalable architecture, consisting of several interconnected components. These components work together to enable seamless data exchange and real-time control of the robot's functionalities.

The architecture of our software is designed to be flexible and adaptable, allowing for easy integration with various devices and systems. It is built using a combination of open-source software and hardware components, ensuring a robust and reliable platform for surveillance tasks.

The key components of our software architecture include:

1. **Web Server**: This component serves as the central hub for data visualization. It provides a user-friendly interface for monitoring and controlling the robot's functionalities.

2. **API**: This component acts as the interface between the web server and the robot's hardware components. It enables real-time data exchange and control of the robot's functionalities.

3. **Mobile App**: This component also serves as the user interface for data visualization. It provides a mobile-friendly experience for monitoring and controlling the robot's functionalities.

4. **Video Streaming**: This component enables real-time video streaming from the robot's cameras. This way we can monitor the robot's surroundings in real-time.

Let's go through each of these components in more detail.

## Web Server

The web server is the backbone of our software architecture. It serves as the central hub for data visualization and provides a user-friendly interface for monitoring and controlling the robot's functionalities.

The web server is built using modern web technologies such as HTML, CSS, and JavaScript. To be more precise, it is built using the following technologies:

- **HTML**: This technology is used to create the structure and layout of the web page.
- **CSS**: This technology is used to style and format the web page.
- **JavaScript**: This technology is used to add interactivity and dynamic behavior to the web page.
- **Node.js**: This technology is used to build the server-side logic of the web server.

The web server is hosted on a Linux server and is accessible through a web browser. That's why one of the requirements for the application is to have a centralized server to host the web server.

## API

Because the API is a important component of our software architecture, we decided to create a separate section for it. You can find more information about the API in the [API](./api.md) section.

## Mobile App

The mobile app is a user interface for data visualization. It provides a mobile-friendly experience for monitoring and controlling the robot's functionalities.

The mobile app is built using a different set of technologies than the web server. It is built using Flutter, a cross-platform mobile application development framework.

Flutter is a popular choice for building mobile applications because it allows developers to write code once and deploy it on multiple platforms, including iOS, Android, and web. This is one of the reasons why we chose Flutter for the mobile app.

## Video Streaming

Video streaming is a crucial feature of our surveillance robot project. It enables real-time video streaming from the robot's cameras. This way we can monitor the robot's surroundings in real-time.

To implement video streaming, we used the `opencv` library, which is a popular open-source computer vision library. It provides a wide range of functionalities for image and video processing, including video streaming.

The only requirement for the video streaming feature is to have a camera connected to the robot. To know how to enable this feature, please refer to the [Installation](./installation.md) section.

