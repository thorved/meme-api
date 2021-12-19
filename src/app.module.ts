import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import {MemeModule} from './meme/meme.module';

@Module({
  imports: [MemeModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
