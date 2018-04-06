import { Controller, Get, Param, Post, Res, UseInterceptors, FileInterceptor, UploadedFile, Req } from '@nestjs/common'

@Controller('package')
export class AppController {

  private packages = { }

  @Get(':pkg')
  getPackage(@Param('pkg') pkg, @Res() res) {
    const [ name, version = 'latest' ] = pkg.split('@')
    console.log(`Fetch: ${name}@${version}`)

    res.set('Content-Type', 'application/zip').end(this.packages[name][version])
  }

  @Post(':pkg')
  @UseInterceptors(FileInterceptor('file'))
  publishPackage(@Param('pkg') pkg, @UploadedFile() file) {
    const [ name, version ] = pkg.split('@')
    console.log(`Publish: ${name}@${version}`)

    const buff = new Buffer(file.buffer, 'binary')

    if (!this.packages[name]) {
      this.packages[name] = { }
    }

    this.packages[name][version] = buff
  }
}
