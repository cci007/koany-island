#!/bin/bash


for FOLDER in $(ls ../ |grep _); do echo $FOLDER; 
cat >> ./config.yml << EOF
      - path-filtering/filter:
          base-revision: main
          name: check-updated-files-${FOLDER}
          mapping: |
            ${FOLDER}/.* run-${FOLDER}-job true
          config-path: ${FOLDER}/.circleci/config.yml

EOF

#cat >> ../${FOLDER}/.circleci/config.yml  << EOF
#
#parameters:
#  run-${FOLDER}-job:
#    type: boolean
#    default: false
#
#
#workflows:
#  ${FOLDER}:
#    when: << pipeline.parameters.run-${FOLDER}-job >>
#    jobs:
#      - build
#
#EOF

done
