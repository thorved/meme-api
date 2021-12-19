import { HttpModule } from '@nestjs/axios';
import { Module } from '@nestjs/common';
// import { UserSchema } from './user.schema';
import { MemeController } from './meme.controller';
import { MemeService } from './meme.service';

@Module({
  imports: [HttpModule],
  controllers: [MemeController],
  providers: [MemeService],
})
export class MemeModule {}
