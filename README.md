This simple project goal was to split lines in an input csv

One of my colleague sent me a file in excel with the following format (CSV):

&lt;id&gt;,&lt;hostname&gt;,&lt;used by&gt;

The latter column contained a comma-delimited list of users.
I need that file to be integrated in a tool that counts the sum of usages. And each usage that is shared among 2 people counts for 0.5, as an equal split

example :

````
id,hostname,users
1,machine1,"herbert,rafaello,jerome"
2,machine2,"jose,jerome"
3,machine3,"allen"
````

The goal in the end is to be able to produce a report :
````
herbert, 0.33333
rafaello, 0.33333
jerome, 0.833333
````
etc.

So I needed a first step that transformed the first csv file into :
````
id,hostname,user,Fraction
1,machine1,herbert,0.3333333
1,machine1,rafaello,0.3333333
1,machine1,jerome,0.3333333
2,machine2,jose,0.5
2,machine2,jerome,0.5
3,machine3,allen,1.0
````
... Then my computation (sum, filter, whatever) can be done, ex pivot table in excel, elastic search + kibana, etc.


