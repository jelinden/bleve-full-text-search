# Testing Bleve Full text search
(http://www.blevesearch.com/)

Full text search is something that you need from time to time and it's not that trivial to implement if you're dataset or queries are complex.

Solr and Elasticsearch are the most used big scale alternatives but they are also taking quite much memory and resources.

Welcome Bleve! Bleve is much more light weight solution than it's rivals. It's done with Go and the database can be chosen. The default is BoltDB, but others can be used too, for example SQlite (https://medium.com/developers-writing/full-text-search-and-indexing-with-bleve-part-1-bd73599d82ef#.xi7gtbinb)

The code tests the most basic operability.
- making the index
- getting all items in index
- getting items by a field value
- sorting
- deleting an item