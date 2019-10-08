#
# Sentinel - Copyright (c) 2019 by www.gatblau.org
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software distributed under
# the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
# either express or implied.
# See the License for the specific language governing permissions and limitations under the License.
#
# Contributors to this project, hereby assign copyright in this code to the project,
# to be licensed under the same terms as the rest of the code.
#
# This vagrant box is used to build OCI images with buildah and CentOS 8
#  to avoid depending on the underlying OS
#
# How to use:
#   $ vagrant up
#   $ vagrant ssh
#   $ cd /app
#   $ make set-version
#   $ make snapshot-image (or make release-image)
#   $ sudo buildah images
#
$script = <<SCRIPT
# install required tools
sudo yum module install -y container-tools
sudo yum install -y git
SCRIPT

Vagrant.configure("2") do |config|
  config.vm.box = "generic/centos8"
  config.vm.box_check_update = true
  config.vm.hostname = "builder"
  config.vm.synced_folder ".", "/app"
  config.vm.provision "shell", inline: $script
end
