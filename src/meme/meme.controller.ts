import {
  Body,
  Headers,
  Controller,
  Delete,
  Get,
  Post,
  Put,
  UseGuards,
  Param,
  StreamableFile,
  Header,
} from '@nestjs/common';
import { createReadStream } from 'fs';
import { join } from 'path';
import { MemeService } from './meme.service';
@Controller('/meme.jpg')
export class MemeController {
  constructor(private readonly memeService: MemeService) {}

  @Get(':subreddit')
  @Header('Cache-Control', 'none',)
  async getMemeSubredit(@Param('subreddit') subreddit: string) {
    console.log(subreddit.split('.')[0]);
    return await this.memeService.getMeme(subreddit.split('.')[0]);
  }
  @Get()
  @Header('Cache-Control', 'none')
  async getMeme(@Param('subreddit') subreddit: string) {
    console.log(subreddit);
    return await this.memeService.getMeme(null);
  }
  //   @Get(':subreddit')
  //   async getMeme(@Param('subreddit') subreddit: string): Promise<StreamableFile> {
  //     await this.memeService.getMeme(subreddit);
  //     //console.log(filename.toString());
  //     const file = createReadStream(join(process.cwd(), "./img/123.jpg"));
  //     return await new StreamableFile(file);
  //   }
}
