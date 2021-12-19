import { HttpService } from '@nestjs/axios';
import { Injectable, StreamableFile } from '@nestjs/common';
import { createReadStream } from 'fs';
import { join } from 'path';

const Fs = require('fs');

const Axios = require('axios');
import { map } from 'rxjs/internal/operators/map';

async function downloadImage(url1: string, filename: string) {
  const url = url1;

  const writer = Fs.createWriteStream('./img/' + filename);

  const response = await Axios({
    url,
    method: 'GET',
    responseType: 'stream',
  });

  response.data.pipe(writer);

  return new Promise((resolve, reject) => {
    writer.on('finish', resolve);
    writer.on('error', reject);
  });
}

function deleteFiles(files, callback) {
  var i = files.length;
  files.forEach(function (filepath) {
    Fs.unlink('img/' + filepath, function (err) {
      i--;
      if (err) {
        callback(err);
        return;
      } else if (i <= 0) {
        callback(null);
      }
    });
  });
}

@Injectable()
export class MemeService {
  constructor(private httpService: HttpService) {}

  async getMeme(subreddit: string) {
    Fs.readdir('./img/', (err, files) => {
      deleteFiles(files, function (err) {
        if (err) {
          console.log(err);
        } else {
          console.log('all files removed');
        }
      });
    });
    if (subreddit === null) {
      var url = `https://meme-api.herokuapp.com/gimme/`;
    } else {
      var url = `https://meme-api.herokuapp.com/gimme/${subreddit}`;
    }

    return this.httpService.get(url).pipe(
      map(async (res) => {
        console.log(res.data['url']);
        var filename = Math.random().toString(16).slice(2).toString();
        return await downloadImage(res.data['url'], filename)
          .then((res) => {
            const file = createReadStream(
              join(process.cwd(), './img/' + filename),
            );

            return new StreamableFile(file);
          })
          .catch((err) => {});
      }),
    );
  }
}
