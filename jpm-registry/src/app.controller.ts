import { Controller, Get, Param, Post, UseInterceptors, FileInterceptor, UploadedFile, Req } from '@nestjs/common'

@Controller('package')
export class AppController {

  @Get(':name')
  getPackage(@Param('name') name) {
    console.log(`Fetch: ${name}`)

    return '#!/env/bin/groovy\n\nprintln \'Hello, World!\''
  }

  @Post(':name')
  publishPackage(@Param('name') name, @Req() req) {
    console.log(`Publish: ${name}`)
    console.log(req.files)
  }
}
