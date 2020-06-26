Blockchain Overview
=====================

Currently the business are running in siloes , Each participants of the business are not connected but integrated so that only enriched information can reach to the target participant. Here all the participants keep their record safe centrally within them to avoid trust and security threats in the future. The record they maintain are mutable in nature and can be change. 
General Blockchain offers the trust between all the participants in the business, 

 Records are hashed within the blocks which maintains the immutability , 

Any change in the record can be tracked on the platforms and can help the participant to track the source of the information.

In blockchain participants are not only integrated but connected and can share information with minimal time and security.

 It is decentralized repository in which all the participants keep their information which avoid the risk of loss of data
 
What is Blockchain
=====================

Blockchain architecture have Participants which we can call as nodes , 
Each node has their own leader and peers where smart contracts executes
And channel where all participants are connected and share information.

when ever the transactions enters the blockchain platform, First the transaction will go to channel and then it will pass to the various participants where their peers will validate the transaction on the basis of the smart contract deployed. Once the transaction meet all the validation parameter then participant will reach to the consensus and creates the block 



How blockchain Works
================

Each block has three key information Current Hash ,Data and Previous hash , Only first block will not have the previous hash and it is called as genesis block.





Loan Industry
============

Banking industry earns most of their revenue by lending money to their customers for  specific time period with rate of interest ,collaterals .

Banks put these attributes in the formula and calculate the return which borrower or the customer of the bank has to pay , Now that payment involves two things principal amount and interest , The interest is the profit of the bank.   

This simple process sounds like really good business.  Now lets move ahead and understand the complexity and the risk involved in this process.

Risk Management 
=============

As per the study published in International Journal of Economic and Business Review , Any bank grants the loan by looking into these pillars Credit Criteria , Gathering Necessary Information about the customer and setting limits on the loan.  Out of three , Credit Criteria and the loan limit , they are dependent on Identification process which involves most of the time for processing the loan. 

Information gathering of the customer involves the customers creditworthiness by looking into its previous loan history and the repayment to other banks or financial institutions , Its collaterals current  value and identification document . These factors are time consuming as all the banks have to collect the information either from the central authorities or they have to ask from the customers and execute the verification process again and again.

Current Loan Process
================

The conventional loan process involves 5 to 6 steps depends on the bank,  Document submission, Document upload, Document verification, eligibility check, decision making but there are some disadvantages to the existing system which have found out from the studies published , 

If in case banks reject the loan then customer has to resubmit the document with other bank and then bank will maintain the information with them and again do the checks and decide.

As all the information stored on bank is centrally located so there will be risk involved to recover the information in case the central location lost the data.

Banks are quite sensitive for their information protection and cyber crime is the big threat to the central system.

Blockchain Solution
==============

As we know that blockchain helps the user to track the asset , In our use case Bank’s Money is an asset which they are lending to the borrowers. So we have created the solution from lender’s view. 

Blockchain known for its decentralization capability can fix the problem of executing identification and verification Process every time with all the banks and minimize the documentation processing

If  lenders agree to join the blockchain network on the same channel then their ledgers will be in sync with all the immutable history of all the customers transactions.

This can help lender verify the borrower and can access its creditworthiness and eventually make their decisions to grant loan in the minimal time and can increase their performance in the market. 

These networks are highly secure and as in peer to peer network their is very low risk of loss of data as the ledgers of all the participants have the same data at the particular point. 


Solution Design 
=============
In our use case we are creating borrower who is looking for a loan and submitted his/her application to the bank and now bank will submit the application, process the information and stores the information on blockchain. 


Technical Design
=============
IN technical design we have these layers… Presentation layer tented the borrower’s Information, Service layer which acts as the middleware layer, security layer which consists of tokens of Blockchain Network , blockchain layer which is storing the information immutably and the database layer which is storing some part of the information which we will see in the POC. 
