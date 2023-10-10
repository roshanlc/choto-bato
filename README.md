### A Basic Link Shortener
> Choto-bato in Nepali means short(quick) route.
* Go (Echo framework)
* Redis


### Endpoints 


<table >
<thead>
<tr>
<th>Method</th>
<th>Endpoint</th>
<th>Request Body</th>
<th>Response</th>
</tr>
</thead>

<tbody>
<tr>
<td>GET</td>
<td>/</td>
<td></td>
<td>Returns this table containing endpoint information.</td>
</tr>

<tr>
<td>POST</td>
<td>/</td>
<td>

{
	"long_url" : "https://somelongurl.com/profile/userX/contentY/welcome"
}

</td>
<td>
{ <br/>
"long_url"  :    "https://somelongurl.com/profile/userX/contentY/welcome",
 <br/>
 "short_url" :	"htttps://choto-bato.com/shortURL"
<br/>
}
</td>
</tr>

<tr>
<td>GET</td>
<td>/shortURL
</td>
<td>
</td>
<td>
{ <br/>
"long_url"  :    "https://somelongurl.com/profile/userX/contentY/welcome",
 <br/>
 "short_url" :	"htttps://choto-bato.com/shortURL"
<br/>
}
</td>
</tr>

</tbody>
</table>

### Running
Requires:
- Go runtime
- Redis

Follow the guide for running a instance.
```bash
git clone github.com/roshanlc/choto-bato

cd choto-bato

# edit redis configuration at store/store.go file

go run main.go
```
