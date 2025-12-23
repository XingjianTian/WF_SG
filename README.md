# WF_SG: Fabric and WireGuard Based Decentralized Smart Grid System

## Overview
This repository implements a decentralized smart grid system integrating HyperLedger Fabric for blockchain consensus and WireGuard for secure E2EE communication. It addresses data security, user privacy, and trustworthy peer transmission in smart grids, supporting modern energy management and renewable sources. The project includes a web application for simulation, visualization, and experiments.

## Abstract
Smart grid has drawn lots of attention and investment these years, which helps not only the modern generation and distribution of traditional power but also highly widens the application of renewable energy sources. However, the key challenges of smart grid are :1. the security of data, the protection of user privacies. 2. the trustful transmission channel between peers. And to help solve these problems, VPN and blockchain can be considered since they have some features perfectly suitable under those situations.

This research introduced WireGuard and HyperLedger Fabric to counter the problems, and proposed a smart grid system integrated with these technologies, in the form of fully implemented web application. Then all the functionalities are displayed and tested, including building a smart device simulator, deploying data visualization and making transactions experiments and WireGuard communication experiments.

Experiment results shows that the bringing of Fabric has certain practical significance and value, which helps ensure the consistency and traceability of transactions. And bringing WireGuard into network infrastructure does not cause too much loss of bandwidth and delay, while ensuring a certain degree of communication security.

## Key Features
- Secure E2EE communication via WireGuard.
- Blockchain-based transaction consistency using HyperLedger Fabric.
- Smart device simulator for testing.
- Data visualization tools.
- Experiments on transactions and communication performance.

## Installation
1. Clone the repo: `git clone https://github.com/XingjianTian/WF_SG.git`
2. Install dependencies: Go for backend, Node.js for web if applicable.
3. Set up HyperLedger Fabric: Follow official docs for network configuration.
4. Configure WireGuard: Generate keys and update conf.yaml.

## Usage
- Run the main.go for backend services.
- Use the web app for visualization and simulations.
- Test transactions and WireGuard communication as per paper experiments.

## License
MIT License (assumed; add LICENSE file if needed).

## Contact
For issues, contact via GitHub or email from the paper.
MIT (assumed; confirm or add LICENSE file).
Contact
For questions, contact the author via GitHub issues.
